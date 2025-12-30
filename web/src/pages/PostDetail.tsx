import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { api } from '../api/client';

interface PostData {
  id: number;
  title: string;
  content: string;
  author_id: number;
  status: string;
  created_at: string;
}

export default function PostDetail() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [post, setPost] = useState<PostData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!id) return;

    const fetchPost = async () => {
      try {
        const data = await api.getPost(parseInt(id, 10));
        setPost(data.post);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to load post');
      } finally {
        setLoading(false);
      }
    };

    fetchPost();
  }, [id]);

  if (loading) {
    return (
      <div className="page">
        <p>Loading...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="page">
        <p className="error">{error}</p>
        <button onClick={() => navigate('/')}>Back to Login</button>
      </div>
    );
  }

  if (!post) {
    return (
      <div className="page">
        <p>Post not found</p>
        <button onClick={() => navigate('/')}>Back to Login</button>
      </div>
    );
  }

  return (
    <div className="page">
      <article className="post-detail">
        <h1>{post.title}</h1>
        <div className="post-meta">
          <span>Author ID: {post.author_id}</span>
          <span>Status: {post.status}</span>
          <span>Created: {new Date(post.created_at).toLocaleString()}</span>
        </div>
        <div className="post-content">
          {post.content}
        </div>
      </article>
      <div className="actions">
        <button onClick={() => navigate('/posts/new')}>Create New Post</button>
      </div>
    </div>
  );
}