import { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import { getCampaigns, deleteCampaign } from '../api/campaigns';
import { Campaign } from '../types/campaign';
import { StatusBadge } from '../components/StatusBadge';
import { LoadingState } from '../components/LoadingState';
import { ErrorMessage } from '../components/ErrorMessage';
import { formatDate } from '../utils/date';

export function CampaignListPage() {
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [deletingId, setDeletingId] = useState<string | null>(null);
  const [deleteError, setDeleteError] = useState<string | null>(null);

  const fetchCampaigns = useCallback(async () => {
    setLoading(true);
    setError(null);
    setDeleteError(null);
    try {
      const data = await getCampaigns();
      setCampaigns(data);
    } catch {
      setError('Failed to load campaigns. Please try again.');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchCampaigns();
  }, [fetchCampaigns]);

  const handleDelete = async (campaign: Campaign) => {
    const confirmed = window.confirm(
      `Are you sure you want to delete "${campaign.title}"?`
    );
    if (!confirmed) return;

    setDeletingId(campaign.id);
    setDeleteError(null);
    try {
      await deleteCampaign(campaign.id);
      await fetchCampaigns();
    } catch {
      setDeleteError(`Failed to delete campaign "${campaign.title}". Please try again.`);
    } finally {
      setDeletingId(null);
    }
  };

  return (
    <div className="page-container">
      <div className="page-header">
        <h1>Campaigns</h1>
        <div className="page-actions">
          <button
            id="btn-refresh-campaigns"
            className="btn btn-secondary"
            onClick={fetchCampaigns}
            disabled={loading}
          >
            Refresh
          </button>
          <Link
            id="link-create-campaign"
            to="/campaigns/new"
            className="btn btn-primary"
          >
            Create Campaign
          </Link>
        </div>
      </div>

      {loading && <LoadingState message="Loading campaigns..." />}

      {!loading && error && <ErrorMessage message={error} />}

      {!loading && deleteError && <ErrorMessage message={deleteError} />}

      {!loading && !error && campaigns.length === 0 && (
        <div className="empty-state">
          <p>No campaigns yet.</p>
          <Link
            id="link-create-first-campaign"
            to="/campaigns/new"
            className="btn btn-primary"
          >
            Create Campaign
          </Link>
        </div>
      )}

      {!loading && !error && campaigns.length > 0 && (
        <div className="table-wrapper">
          <table className="campaign-table">
            <thead>
              <tr>
                <th>Title</th>
                <th>Currency</th>
                <th>Initial Budget</th>
                <th>Remaining Budget</th>
                <th>Impressions</th>
                <th>Status</th>
                <th>Start Date</th>
                <th>End Date</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {campaigns.map((campaign) => (
                <tr key={campaign.id}>
                  <td>{campaign.title}</td>
                  <td>{campaign.currency}</td>
                  <td>{campaign.initial_budget.toFixed(2)}</td>
                  <td>{campaign.remaining_budget.toFixed(2)}</td>
                  <td>{campaign.impression_count}</td>
                  <td>
                    <StatusBadge status={campaign.status} />
                  </td>
                  <td>{formatDate(campaign.start_date)}</td>
                  <td>{formatDate(campaign.end_date)}</td>
                  <td className="action-cell">
                    <Link
                      id={`link-view-${campaign.id}`}
                      to={`/campaigns/${campaign.id}`}
                      className="btn btn-sm btn-secondary"
                    >
                      View Details
                    </Link>
                    <button
                      id={`btn-delete-${campaign.id}`}
                      className="btn btn-sm btn-danger"
                      onClick={() => handleDelete(campaign)}
                      disabled={deletingId === campaign.id}
                    >
                      {deletingId === campaign.id ? 'Deleting...' : 'Delete'}
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
