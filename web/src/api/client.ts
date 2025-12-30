const BASE_URL = '';

export async function apiFetch<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const token = localStorage.getItem('token');

  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  if (token) {
    (headers as Record<string, string>)['Authorization'] = `Bearer ${token}`;
  }

  const response = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: 'Request failed' }));
    throw new Error(error.error || `HTTP ${response.status}`);
  }

  return response.json();
}

export const api = {
  login: (email: string, password: string) =>
    apiFetch<{ token: string; user: { id: number; username: string; email: string; role: string } }>(
      '/auth/login',
      {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      }
    ),

  createPost: (title: string, content: string, status: string) =>
    apiFetch<{ post: { id: number }; message: string }>('/api/posts', {
      method: 'POST',
      body: JSON.stringify({ title, content, status }),
    }),

  getPost: (id: number) =>
    apiFetch<{ post: { id: number; title: string; content: string; author_id: number; status: string; created_at: string } }>(
      `/api/posts/${id}`
    ),
};