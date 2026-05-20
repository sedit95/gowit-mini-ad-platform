import { useState, useEffect, useCallback } from 'react';
import { useParams, Link, useNavigate } from 'react-router-dom';
import {
  getCampaign,
  getStats,
  recordImpression,
  updateCampaign,
  deleteCampaign,
} from '../api/campaigns';
import type {
  Campaign,
  StatsResponse,
  CampaignStatus,
  UpdateCampaignRequest,
  ImpressionResponse,
} from '../types/campaign';
import { StatusBadge } from '../components/StatusBadge';
import { LoadingState } from '../components/LoadingState';
import { ErrorMessage } from '../components/ErrorMessage';
import {
  formatDateTime,
  formatDate,
  toISOStringFromDateTimeLocal,
  toDateTimeLocalValue,
} from '../utils/date';

const VALID_STATUSES: CampaignStatus[] = ['active', 'paused', 'completed'];

interface UpdateFormState {
  title: string;
  currency: string;
  start_date: string;
  end_date: string;
  status: CampaignStatus;
}

interface UpdateFormErrors {
  title?: string;
  currency?: string;
  start_date?: string;
  end_date?: string;
  status?: string;
}

function buildUpdateForm(campaign: Campaign): UpdateFormState {
  return {
    title: campaign.title,
    currency: campaign.currency,
    start_date: toDateTimeLocalValue(campaign.start_date),
    end_date: toDateTimeLocalValue(campaign.end_date),
    status: campaign.status,
  };
}

function validateUpdateForm(form: UpdateFormState): UpdateFormErrors {
  const errors: UpdateFormErrors = {};

  if (!form.title.trim()) {
    errors.title = 'Title is required.';
  }

  const upperCurrency = form.currency.toUpperCase();
  if (!upperCurrency || upperCurrency.length !== 3) {
    errors.currency = 'Currency must be exactly 3 characters (e.g. USD).';
  }

  if (!form.start_date) {
    errors.start_date = 'Start date is required.';
  }

  if (!form.end_date) {
    errors.end_date = 'End date is required.';
  }

  if (form.start_date && form.end_date && form.end_date < form.start_date) {
    errors.end_date = 'End date must be on or after start date.';
  }

  if (!VALID_STATUSES.includes(form.status)) {
    errors.status = 'Invalid status value.';
  }

  return errors;
}

export function CampaignDetailPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  // ── Core State ──────────────────────────────────────────────────────────────
  const [campaign, setCampaign] = useState<Campaign | null>(null);
  const [stats, setStats] = useState<StatsResponse | null>(null);
  const [loadError, setLoadError] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(true);

  // ── Stats Refresh State ──────────────────────────────────────────────────────
  const [statsError, setStatsError] = useState<string>('');
  const [isRefreshingStats, setIsRefreshingStats] = useState<boolean>(false);

  // ── Impression State ─────────────────────────────────────────────────────────
  const [impressionResult, setImpressionResult] = useState<ImpressionResponse | null>(null);
  const [impressionMessage, setImpressionMessage] = useState<string>('');
  const [impressionError, setImpressionError] = useState<string>('');
  const [isRecording, setIsRecording] = useState<boolean>(false);

  // ── Delete State ─────────────────────────────────────────────────────────────
  const [deleteError, setDeleteError] = useState<string>('');
  const [isDeleting, setIsDeleting] = useState<boolean>(false);

  // ── Update Form State ─────────────────────────────────────────────────────────
  const [updateForm, setUpdateForm] = useState<UpdateFormState>({
    title: '',
    currency: '',
    start_date: '',
    end_date: '',
    status: 'active',
  });
  const [updateErrors, setUpdateErrors] = useState<UpdateFormErrors>({});
  const [updateSuccess, setUpdateSuccess] = useState<string>('');
  const [updateError, setUpdateError] = useState<string>('');
  const [isUpdating, setIsUpdating] = useState<boolean>(false);

  // ── Stats fetch ──────────────────────────────────────────────────────────────
  const fetchStats = useCallback(async (campaignId: string): Promise<void> => {
    try {
      const data = await getStats(campaignId);
      setStats(data);
      setStatsError('');
    } catch {
      setStatsError('Failed to load stats.');
    }
  }, []);

  // ── Initial load ─────────────────────────────────────────────────────────────
  useEffect(() => {
    if (!id) {
      setLoadError('Campaign ID is missing from the URL.');
      setIsLoading(false);
      return;
    }

    let cancelled = false;

    async function load() {
      setIsLoading(true);
      setLoadError('');

      try {
        const [campaignData, statsData] = await Promise.all([
          getCampaign(id as string),
          getStats(id as string),
        ]);

        if (!cancelled) {
          setCampaign(campaignData);
          setStats(statsData);
          setUpdateForm(buildUpdateForm(campaignData));
        }
      } catch {
        if (!cancelled) {
          setLoadError('Failed to load campaign. It may not exist or the server is unavailable.');
        }
      } finally {
        if (!cancelled) {
          setIsLoading(false);
        }
      }
    }

    load();

    return () => {
      cancelled = true;
    };
  }, [id]);

  // ── Refresh Stats ─────────────────────────────────────────────────────────────
  async function handleRefreshStats() {
    if (!id) return;
    setIsRefreshingStats(true);
    setStatsError('');
    try {
      await fetchStats(id);
    } finally {
      setIsRefreshingStats(false);
    }
  }

  // ── Record Impression ─────────────────────────────────────────────────────────
  async function handleRecordImpression() {
    if (!id) return;
    setIsRecording(true);
    setImpressionMessage('');
    setImpressionError('');
    setImpressionResult(null);

    try {
      const result = await recordImpression(id);
      setImpressionResult(result);

      if (result.accepted) {
        setImpressionMessage('Impression recorded successfully.');
      } else if (result.reason === 'budget_exhausted') {
        setImpressionMessage('Budget exhausted. No more impressions can be recorded.');
      } else if (result.reason === 'campaign_not_active') {
        setImpressionMessage('Campaign is not active.');
      } else {
        setImpressionMessage('Impression not accepted.');
      }

      // Always refresh stats after recording impression
      await fetchStats(id);
    } catch {
      setImpressionError('Failed to record impression. Please try again.');
    } finally {
      setIsRecording(false);
    }
  }

  // ── Delete Campaign ───────────────────────────────────────────────────────────
  async function handleDelete() {
    if (!id) return;
    const confirmed = window.confirm(
      'Are you sure you want to delete this campaign? This action cannot be undone.'
    );
    if (!confirmed) return;

    setIsDeleting(true);
    setDeleteError('');

    try {
      await deleteCampaign(id);
      navigate('/');
    } catch {
      setDeleteError('Failed to delete campaign. Please try again.');
      setIsDeleting(false);
    }
  }

  // ── Update Form Handlers ──────────────────────────────────────────────────────
  function handleUpdateFieldChange(
    field: keyof UpdateFormState,
    value: string
  ) {
    setUpdateForm((prev) => ({ ...prev, [field]: value }));
    if (updateErrors[field]) {
      setUpdateErrors((prev) => ({ ...prev, [field]: undefined }));
    }
    setUpdateSuccess('');
    setUpdateError('');
  }

  async function handleUpdate(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!id) return;

    const errors = validateUpdateForm(updateForm);
    if (Object.keys(errors).length > 0) {
      setUpdateErrors(errors);
      return;
    }

    setIsUpdating(true);
    setUpdateSuccess('');
    setUpdateError('');

    const payload: UpdateCampaignRequest = {
      title: updateForm.title.trim(),
      currency: updateForm.currency.toUpperCase(),
      start_date: toISOStringFromDateTimeLocal(updateForm.start_date),
      end_date: toISOStringFromDateTimeLocal(updateForm.end_date),
      status: updateForm.status,
    };

    try {
      const updated = await updateCampaign(id, payload);
      setCampaign(updated);
      setUpdateForm(buildUpdateForm(updated));
      setUpdateErrors({});
      setUpdateSuccess('Campaign updated successfully.');
      await fetchStats(id);
    } catch {
      setUpdateError('Failed to update campaign. Please check your input and try again.');
    } finally {
      setIsUpdating(false);
    }
  }

  // ── Derived helpers ───────────────────────────────────────────────────────────
  const isImpressionDisabled =
    isRecording ||
    (stats !== null &&
      (stats.status !== 'active' || stats.remaining_budget <= 0));

  // ── Render ────────────────────────────────────────────────────────────────────
  if (isLoading) {
    return (
      <div className="page-container">
        <LoadingState />
      </div>
    );
  }

  if (loadError || !campaign) {
    return (
      <div className="page-container">
        <div className="page-header">
          <Link to="/" className="btn btn-secondary btn-sm">
            ← Back to Campaigns
          </Link>
        </div>
        <ErrorMessage message={loadError || 'Campaign not found.'} />
      </div>
    );
  }

  return (
    <div className="page-container">
      {/* ── Header ── */}
      <div className="page-header">
        <h1 className="detail-page-title">{campaign.title}</h1>
        <div className="page-actions">
          <Link to="/" className="btn btn-secondary btn-sm">
            ← Back
          </Link>
          <button
            id="btn-delete-campaign"
            className="btn btn-danger btn-sm"
            onClick={handleDelete}
            disabled={isDeleting}
          >
            {isDeleting ? 'Deleting…' : 'Delete Campaign'}
          </button>
        </div>
      </div>

      {deleteError && <ErrorMessage message={deleteError} />}

      {/* ── Campaign Detail Card ── */}
      <section className="detail-section" aria-label="Campaign Details">
        <h2 className="detail-section-title">Campaign Details</h2>
        <div className="detail-card">
          <dl className="detail-grid">
            <div className="detail-item">
              <dt className="detail-label">Title</dt>
              <dd className="detail-value">{campaign.title}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Status</dt>
              <dd className="detail-value">
                <StatusBadge status={campaign.status} />
              </dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Currency</dt>
              <dd className="detail-value">{campaign.currency}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Initial Budget</dt>
              <dd className="detail-value">
                {campaign.initial_budget.toFixed(2)} {campaign.currency}
              </dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Remaining Budget</dt>
              <dd className="detail-value">
                {campaign.remaining_budget.toFixed(2)} {campaign.currency}
              </dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Impression Count</dt>
              <dd className="detail-value">{campaign.impression_count}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Start Date</dt>
              <dd className="detail-value">{formatDate(campaign.start_date)}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">End Date</dt>
              <dd className="detail-value">{formatDate(campaign.end_date)}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Created At</dt>
              <dd className="detail-value">{formatDateTime(campaign.created_at)}</dd>
            </div>
            <div className="detail-item">
              <dt className="detail-label">Updated At</dt>
              <dd className="detail-value">{formatDateTime(campaign.updated_at)}</dd>
            </div>
          </dl>
        </div>
      </section>

      {/* ── Stats Section ── */}
      <section className="detail-section" aria-label="Campaign Stats">
        <div className="detail-section-header">
          <h2 className="detail-section-title">Stats</h2>
          <button
            id="btn-refresh-stats"
            className="btn btn-secondary btn-sm"
            onClick={handleRefreshStats}
            disabled={isRefreshingStats}
          >
            {isRefreshingStats ? 'Refreshing…' : 'Refresh Stats'}
          </button>
        </div>

        {statsError && <ErrorMessage message={statsError} />}

        {stats && (
          <div className="stats-grid">
            <div className="stat-card">
              <span className="stat-label">Total Impressions</span>
              <span className="stat-value">{stats.total_impressions}</span>
            </div>
            <div className="stat-card">
              <span className="stat-label">Initial Budget</span>
              <span className="stat-value">
                {stats.initial_budget.toFixed(2)} {stats.currency}
              </span>
            </div>
            <div className="stat-card">
              <span className="stat-label">Spent Budget</span>
              <span className="stat-value">
                {stats.spent_budget.toFixed(2)} {stats.currency}
              </span>
            </div>
            <div className="stat-card">
              <span className="stat-label">Remaining Budget</span>
              <span className="stat-value">
                {stats.remaining_budget.toFixed(2)} {stats.currency}
              </span>
            </div>
            <div className="stat-card">
              <span className="stat-label">Status</span>
              <span className="stat-value">
                <StatusBadge status={stats.status} />
              </span>
            </div>
          </div>
        )}
      </section>

      {/* ── Record Impression Section ── */}
      <section className="detail-section" aria-label="Record Impression">
        <h2 className="detail-section-title">Record Impression</h2>

        <div className="impression-action-row">
          <button
            id="btn-record-impression"
            className="btn btn-primary"
            onClick={handleRecordImpression}
            disabled={isImpressionDisabled}
          >
            {isRecording ? 'Recording…' : 'Record Impression'}
          </button>
        </div>

        {impressionError && <ErrorMessage message={impressionError} />}

        {impressionMessage && (
          <div
            className={
              impressionResult?.accepted
                ? 'impression-message impression-message--success'
                : 'impression-message impression-message--info'
            }
            role="status"
            aria-live="polite"
          >
            <p>{impressionMessage}</p>
            {impressionResult && (
              <p className="impression-meta">
                Remaining budget: {impressionResult.remaining_budget.toFixed(2)}{' '}
                {campaign.currency}
                {impressionResult.impression_count !== undefined && (
                  <> · Impressions: {impressionResult.impression_count}</>
                )}
              </p>
            )}
          </div>
        )}
      </section>

      {/* ── Update Section ── */}
      <section className="detail-section" aria-label="Update Campaign">
        <h2 className="detail-section-title">Update Campaign</h2>

        {updateSuccess && (
          <div className="update-success" role="status" aria-live="polite">
            <p>{updateSuccess}</p>
          </div>
        )}
        {updateError && <ErrorMessage message={updateError} />}

        <form
          id="form-update-campaign"
          className="form-card detail-update-form"
          onSubmit={handleUpdate}
          noValidate
        >
          {/* Title */}
          <div className="form-group">
            <label htmlFor="update-title" className="form-label">
              Title <span aria-hidden="true">*</span>
            </label>
            <input
              id="update-title"
              type="text"
              className={`form-input${updateErrors.title ? ' form-input-error' : ''}`}
              value={updateForm.title}
              onChange={(e) => handleUpdateFieldChange('title', e.target.value)}
              disabled={isUpdating}
              maxLength={200}
            />
            {updateErrors.title && (
              <p className="field-error" role="alert">
                {updateErrors.title}
              </p>
            )}
          </div>

          {/* Currency */}
          <div className="form-group">
            <label htmlFor="update-currency" className="form-label">
              Currency <span aria-hidden="true">*</span>
            </label>
            <input
              id="update-currency"
              type="text"
              className={`form-input${updateErrors.currency ? ' form-input-error' : ''}`}
              value={updateForm.currency}
              onChange={(e) =>
                handleUpdateFieldChange('currency', e.target.value.toUpperCase())
              }
              disabled={isUpdating}
              maxLength={3}
              placeholder="USD"
            />
            {updateErrors.currency && (
              <p className="field-error" role="alert">
                {updateErrors.currency}
              </p>
            )}
          </div>

          {/* Start Date */}
          <div className="form-group">
            <label htmlFor="update-start-date" className="form-label">
              Start Date <span aria-hidden="true">*</span>
            </label>
            <input
              id="update-start-date"
              type="datetime-local"
              className={`form-input${updateErrors.start_date ? ' form-input-error' : ''}`}
              value={updateForm.start_date}
              onChange={(e) => handleUpdateFieldChange('start_date', e.target.value)}
              disabled={isUpdating}
            />
            {updateErrors.start_date && (
              <p className="field-error" role="alert">
                {updateErrors.start_date}
              </p>
            )}
          </div>

          {/* End Date */}
          <div className="form-group">
            <label htmlFor="update-end-date" className="form-label">
              End Date <span aria-hidden="true">*</span>
            </label>
            <input
              id="update-end-date"
              type="datetime-local"
              className={`form-input${updateErrors.end_date ? ' form-input-error' : ''}`}
              value={updateForm.end_date}
              onChange={(e) => handleUpdateFieldChange('end_date', e.target.value)}
              disabled={isUpdating}
            />
            {updateErrors.end_date && (
              <p className="field-error" role="alert">
                {updateErrors.end_date}
              </p>
            )}
          </div>

          {/* Status */}
          <div className="form-group">
            <label htmlFor="update-status" className="form-label">
              Status <span aria-hidden="true">*</span>
            </label>
            <select
              id="update-status"
              className={`form-input${updateErrors.status ? ' form-input-error' : ''}`}
              value={updateForm.status}
              onChange={(e) =>
                handleUpdateFieldChange('status', e.target.value as CampaignStatus)
              }
              disabled={isUpdating}
            >
              <option value="active">Active</option>
              <option value="paused">Paused</option>
              <option value="completed">Completed</option>
            </select>
            {updateErrors.status && (
              <p className="field-error" role="alert">
                {updateErrors.status}
              </p>
            )}
          </div>

          <div className="form-actions">
            <button
              id="btn-update-campaign"
              type="submit"
              className="btn btn-primary"
              disabled={isUpdating}
            >
              {isUpdating ? 'Saving…' : 'Save Changes'}
            </button>
          </div>
        </form>
      </section>
    </div>
  );
}
