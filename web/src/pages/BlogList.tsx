import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { api, Post } from '../api/client';

export default function BlogList() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const data = await api.getPublishedPosts();
        setPosts(data.posts || []);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to load posts');
      } finally {
        setLoading(false);
      }
    };
    fetchPosts();
  }, []);

  if (loading) {
    return <div className="page"><p>Loading...</p></div>;
  }

  if (error) {
    return <div className="page"><p className="error">{error}</p></div>;
  }

  return (
    <div className="page">
      <header className="blog-header">
        <h1>Blog</h1>
        <Link to="/login" className="login-link">Admin Login</Link>
      </header>

      {posts.length === 0 ? (
        <p className="empty-message">No posts yet.</p>
      ) : (
        <div className="post-list">
          {posts.map((post) => (
            <article key={post.id} className="post-card">
              <Link to={`/posts/${post.id}`}>
                <h2>{post.title}</h2>
              </Link>
              <p className="post-excerpt">
                {post.content.length > 150
                  ? post.content.substring(0, 150) + '...'
                  : post.content}
              </p>
              <time className="post-date">
                {new Date(post.created_at).toLocaleDateString()}
              </time>
            </article>
          ))}
        </div>
      )}
    </div>
  );
}