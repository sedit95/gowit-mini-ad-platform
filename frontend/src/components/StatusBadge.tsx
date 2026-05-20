import { CampaignStatus } from '../types/campaign';

interface StatusBadgeProps {
  status: CampaignStatus;
}

export function StatusBadge({ status }: StatusBadgeProps) {
  let label = '';
  switch (status) {
    case 'active':
      label = 'Active';
      break;
    case 'paused':
      label = 'Paused';
      break;
    case 'completed':
      label = 'Completed';
      break;
    default:
      label = status;
  }

  return (
    <span className={`status-badge status-${status}`}>
      {label}
    </span>
  );
}
