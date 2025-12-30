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

export interface Post {
  id: number;
  title: string;
  content: string;
  author_id: number;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
  role: string;
}

export const api = {
  login: (email: string, password: string) =>
    apiFetch<{ token: string; user: User; message: string }>(
      '/auth/login',
      {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      }
    ),

  // Public endpoints
  getPublishedPosts: () =>
    apiFetch<{ posts: Post[] | null }>('/api/posts'),

  getPost: (id: number) =>
    apiFetch<{ post: Post }>(`/api/posts/${id}`),

  // Protected endpoints (require auth)
  getAllPosts: () =>
    apiFetch<{ posts: Post[] | null }>('/api/posts'),

  createPost: (title: string, content: string, status: string) =>
    apiFetch<{ post: Post; message: string }>('/api/posts', {
      method: 'POST',
      body: JSON.stringify({ title, content, status }),
    }),

  updatePost: (id: number, data: { title?: string; content?: string; status?: string }) =>
    apiFetch<{ post: Post; message: string }>(`/api/posts/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    }),

  deletePost: (id: number) =>
    apiFetch<{ message: string }>(`/api/posts/${id}`, {
      method: 'DELETE',
    }),
};