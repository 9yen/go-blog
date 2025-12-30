import { useState, useEffect } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
import { api, Post } from '../api/client';

export default function PostDetail() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [post, setPost] = useState<Post | null>(null);
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
    return <div className="page"><p>Loading...</p></div>;
  }

  if (error) {
    return (
      <div className="page">
        <p className="error">{error}</p>
        <button onClick={() => navigate('/')}>Back to Blog</button>
      </div>
    );
  }

  if (!post) {
    return (
      <div className="page">
        <p>Post not found</p>
        <button onClick={() => navigate('/')}>Back to Blog</button>
      </div>
    );
  }

  return (
    <div className="page">
      <nav className="breadcrumb">
        <Link to="/">← Back to Blog</Link>
      </nav>
      <article className="post-detail">
        <h1>{post.title}</h1>
        <div className="post-meta">
          <time>{new Date(post.created_at).toLocaleDateString()}</time>
        </div>
        <div className="post-content">
          {post.content}
        </div>
      </article>
    </div>
  );
}