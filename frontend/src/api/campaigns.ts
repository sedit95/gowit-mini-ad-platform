import {
  Campaign,
  CreateCampaignRequest,
  UpdateCampaignRequest,
  StatsResponse,
  ImpressionResponse,
} from '../types/campaign';
import { apiRequest } from './client';

export async function getCampaigns(): Promise<Campaign[]> {
  return apiRequest<Campaign[]>('/campaigns', {
    method: 'GET',
  });
}

export async function createCampaign(payload: CreateCampaignRequest): Promise<Campaign> {
  return apiRequest<Campaign>('/campaigns', {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export async function getCampaign(id: string): Promise<Campaign> {
  return apiRequest<Campaign>(`/campaigns/${encodeURIComponent(id)}`, {
    method: 'GET',
  });
}

export async function updateCampaign(id: string, payload: UpdateCampaignRequest): Promise<Campaign> {
  return apiRequest<Campaign>(`/campaigns/${encodeURIComponent(id)}`, {
    method: 'PUT',
    body: JSON.stringify(payload),
  });
}

export async function deleteCampaign(id: string): Promise<void> {
  return apiRequest<void>(`/campaigns/${encodeURIComponent(id)}`, {
    method: 'DELETE',
  });
}

export async function recordImpression(id: string): Promise<ImpressionResponse> {
  return apiRequest<ImpressionResponse>(`/impression/${encodeURIComponent(id)}`, {
    method: 'POST',
  });
}

export async function getStats(id: string): Promise<StatsResponse> {
  return apiRequest<StatsResponse>(`/stats/${encodeURIComponent(id)}`, {
    method: 'GET',
  });
}
