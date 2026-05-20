import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { createCampaign } from '../api/campaigns';
import { CampaignStatus, CreateCampaignRequest } from '../types/campaign';
import { ErrorMessage } from '../components/ErrorMessage';
import { toISOStringFromDateTimeLocal } from '../utils/date';

interface FormState {
  title: string;
  budget: string;
  currency: string;
  start_date: string;
  end_date: string;
  status: CampaignStatus;
}

interface FormErrors {
  title?: string;
  budget?: string;
  currency?: string;
  start_date?: string;
  end_date?: string;
  status?: string;
}

const VALID_STATUSES: CampaignStatus[] = ['active', 'paused', 'completed'];

function validate(form: FormState): FormErrors {
  const errors: FormErrors = {};

  if (!form.title.trim()) {
    errors.title = 'Title is required.';
  }

  if (!form.budget.trim()) {
    errors.budget = 'Budget is required.';
  } else {
    const budgetNum = Number(form.budget);
    if (!Number.isInteger(budgetNum) || budgetNum <= 0) {
      errors.budget = 'Budget must be a whole number greater than 0.';
    }
  }

  const currency = form.currency.trim().toUpperCase();
  if (!currency) {
    errors.currency = 'Currency is required.';
  } else if (currency.length !== 3) {
    errors.currency = 'Currency must be exactly 3 characters (e.g. USD).';
  }

  if (!form.start_date) {
    errors.start_date = 'Start date is required.';
  }

  if (!form.end_date) {
    errors.end_date = 'End date is required.';
  }

  if (form.start_date && form.end_date && form.end_date < form.start_date) {
    errors.end_date = 'End date must be greater than or equal to start date.';
  }

  if (!VALID_STATUSES.includes(form.status)) {
    errors.status = 'Status must be active, paused, or completed.';
  }

  return errors;
}

export function CampaignCreatePage() {
  const navigate = useNavigate();

  const [form, setForm] = useState<FormState>({
    title: '',
    budget: '',
    currency: '',
    start_date: '',
    end_date: '',
    status: 'active',
  });

  const [errors, setErrors] = useState<FormErrors>({});
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
    // Clear field error on change
    setErrors((prev) => ({ ...prev, [name]: undefined }));
    setApiError(null);
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const validationErrors = validate(form);
    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }

    const startISO = toISOStringFromDateTimeLocal(form.start_date);
    const endISO = toISOStringFromDateTimeLocal(form.end_date);

    if (!startISO || !endISO) {
      setErrors({
        start_date: !startISO ? 'Invalid start date.' : undefined,
        end_date: !endISO ? 'Invalid end date.' : undefined,
      });
      return;
    }

    const payload: CreateCampaignRequest = {
      title: form.title.trim(),
      budget: Number(form.budget),
      currency: form.currency.trim().toUpperCase(),
      start_date: startISO,
      end_date: endISO,
      status: form.status,
    };

    setSubmitting(true);
    setApiError(null);

    try {
      const created = await createCampaign(payload);
      navigate(`/campaigns/${created.id}`);
    } catch (err: unknown) {
      const message =
        err instanceof Error
          ? err.message
          : 'Failed to create campaign. Please try again.';
      setApiError(message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="page-container">
      <div className="page-header">
        <h1>Create Campaign</h1>
        <Link
          id="link-back-to-list"
          to="/"
          className="btn btn-secondary"
        >
          Back to List
        </Link>
      </div>

      {apiError && <ErrorMessage message={apiError} />}

      <div className="form-card">
        <form
          id="form-create-campaign"
          onSubmit={handleSubmit}
          noValidate
        >
          {/* Title */}
          <div className="form-group">
            <label htmlFor="field-title" className="form-label">
              Title
            </label>
            <input
              id="field-title"
              name="title"
              type="text"
              className={`form-input${errors.title ? ' form-input-error' : ''}`}
              value={form.title}
              onChange={handleChange}
              placeholder="Campaign title"
              disabled={submitting}
            />
            {errors.title && (
              <p className="field-error">{errors.title}</p>
            )}
          </div>

          {/* Budget */}
          <div className="form-group">
            <label htmlFor="field-budget" className="form-label">
              Budget
            </label>
            <input
              id="field-budget"
              name="budget"
              type="number"
              min="1"
              step="1"
              className={`form-input${errors.budget ? ' form-input-error' : ''}`}
              value={form.budget}
              onChange={handleChange}
              placeholder="e.g. 1000"
              disabled={submitting}
            />
            {errors.budget && (
              <p className="field-error">{errors.budget}</p>
            )}
          </div>

          {/* Currency */}
          <div className="form-group">
            <label htmlFor="field-currency" className="form-label">
              Currency
            </label>
            <input
              id="field-currency"
              name="currency"
              type="text"
              maxLength={3}
              className={`form-input${errors.currency ? ' form-input-error' : ''}`}
              value={form.currency}
              onChange={handleChange}
              placeholder="e.g. USD"
              disabled={submitting}
            />
            {errors.currency && (
              <p className="field-error">{errors.currency}</p>
            )}
          </div>

          {/* Start Date */}
          <div className="form-group">
            <label htmlFor="field-start-date" className="form-label">
              Start Date
            </label>
            <input
              id="field-start-date"
              name="start_date"
              type="datetime-local"
              className={`form-input${errors.start_date ? ' form-input-error' : ''}`}
              value={form.start_date}
              onChange={handleChange}
              disabled={submitting}
            />
            {errors.start_date && (
              <p className="field-error">{errors.start_date}</p>
            )}
          </div>

          {/* End Date */}
          <div className="form-group">
            <label htmlFor="field-end-date" className="form-label">
              End Date
            </label>
            <input
              id="field-end-date"
              name="end_date"
              type="datetime-local"
              className={`form-input${errors.end_date ? ' form-input-error' : ''}`}
              value={form.end_date}
              onChange={handleChange}
              disabled={submitting}
            />
            {errors.end_date && (
              <p className="field-error">{errors.end_date}</p>
            )}
          </div>

          {/* Status */}
          <div className="form-group">
            <label htmlFor="field-status" className="form-label">
              Status
            </label>
            <select
              id="field-status"
              name="status"
              className={`form-input${errors.status ? ' form-input-error' : ''}`}
              value={form.status}
              onChange={handleChange}
              disabled={submitting}
            >
              <option value="active">Active</option>
              <option value="paused">Paused</option>
              <option value="completed">Completed</option>
            </select>
            {errors.status && (
              <p className="field-error">{errors.status}</p>
            )}
          </div>

          {/* Submit */}
          <div className="form-actions">
            <button
              id="btn-submit-create"
              type="submit"
              className="btn btn-primary"
              disabled={submitting}
            >
              {submitting ? 'Creating...' : 'Create Campaign'}
            </button>
            <Link
              id="link-cancel-create"
              to="/"
              className="btn btn-secondary"
            >
              Cancel
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
