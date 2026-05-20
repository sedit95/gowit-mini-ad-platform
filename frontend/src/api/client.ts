import { ApiError } from '../types/campaign';

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export async function apiRequest<T>(path: string, options?: RequestInit): Promise<T> {
  const url = `${BASE_URL}${path}`;

  const headers = new Headers(options?.headers);
  if (options?.body && !headers.has('Content-Type')) {
    headers.set('Content-Type', 'application/json');
  }

  const config: RequestInit = {
    ...options,
    headers,
  };

  let response: Response;
  try {
    response = await fetch(url, config);
  } catch (error) {
    const err: ApiError = {
      status: 0,
      code: 'network_error',
      message: 'Failed to connect to the server. Please check your internet connection.',
    };
    throw err;
  }

  // Handle 204 No Content
  if (response.status === 204) {
    return undefined as unknown as T;
  }

  // Parse JSON response body
  let data: any;
  try {
    const text = await response.text();
    data = text ? JSON.parse(text) : {};
  } catch (error) {
    const err: ApiError = {
      status: response.status,
      code: 'parse_error',
      message: 'Failed to parse JSON response from the server.',
    };
    throw err;
  }

  if (!response.ok) {
    // Parse backend error envelope: { "error": { "code": "...", "message": "..." } }
    let code = 'unknown_error';
    let message = 'An unexpected error occurred.';

    if (data && data.error) {
      if (typeof data.error.code === 'string') {
        code = data.error.code;
      }
      if (typeof data.error.message === 'string') {
        message = data.error.message;
      }
    } else if (data && typeof data.message === 'string') {
        message = data.message;
    }

    const err: ApiError = {
      status: response.status,
      code,
      message,
    };
    throw err;
  }

  return data as T;
}
