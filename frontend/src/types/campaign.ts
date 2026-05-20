export type CampaignStatus = "active" | "paused" | "completed";

export interface Campaign {
  id: string;
  title: string;
  currency: string;
  initial_budget: number;
  remaining_budget: number;
  impression_count: number;
  status: CampaignStatus;
  start_date: string;
  end_date: string;
  created_at: string;
  updated_at: string;
}

export interface CreateCampaignRequest {
  title: string;
  budget: number;
  currency: string;
  start_date: string;
  end_date: string;
  status?: CampaignStatus;
}

export interface UpdateCampaignRequest {
  title: string;
  currency: string;
  start_date: string;
  end_date: string;
  status: CampaignStatus;
}

export interface StatsResponse {
  campaign_id: string;
  title: string;
  currency: string;
  total_impressions: number;
  initial_budget: number;
  spent_budget: number;
  remaining_budget: number;
  status: CampaignStatus;
}

export type ImpressionFailureReason = "budget_exhausted" | "campaign_not_active";

export interface ImpressionResponse {
  campaign_id: string;
  accepted: boolean;
  reason?: ImpressionFailureReason;
  remaining_budget: number;
  impression_count?: number;
  status: CampaignStatus;
}

export type ApiErrorCode =
  | "validation_error"
  | "not_found"
  | "invalid_status"
  | "internal_error"
  | "budget_exhausted"
  | "campaign_not_active";

export interface ApiError {
  status: number;
  code: ApiErrorCode | string;
  message: string;
}
