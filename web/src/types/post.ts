export interface User {
  id: number;
  username: string;
  email: string;
  role: string;
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

export interface LoginResponse {
  token: string;
  user: User;
}

export interface PostResponse {
  post: Post;
  message?: string;
}