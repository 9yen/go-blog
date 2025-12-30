import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { api, Post } from '../../api/client';
import { removeToken } from '../../utils/auth';

export default function Dashboard() {
  const navigate = useNavigate();
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  const fetchPosts = async () => {
    try {
      const data = await api.getAllPosts();
      setPosts(data.posts || []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load posts');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  const handleDelete = async (id: number) => {
    if (!confirm('Are you sure you want to delete this post?')) return;
    
    try {
      await api.deletePost(id);
      setPosts(posts.filter(p => p.id !== id));
    } catch (err) {
      alert(err instanceof Error ? err.message : 'Failed to delete post');
    }
  };

  const handleLogout = () => {
    removeToken();
    navigate('/login');
  };

  if (loading) {
    return <div className="page"><p>Loading...</p></div>;
  }

  return (
    <div className="page admin-page">
      <header className="admin-header">
        <h1>Admin Dashboard</h1>
        <div className="admin-actions">
          <Link to="/" className="btn-secondary">View Blog</Link>
          <button onClick={handleLogout} className="btn-secondary">Logout</button>
        </div>
      </header>

      {error && <p className="error">{error}</p>}

      <div className="admin-toolbar">
        <Link to="/admin/posts/new" className="btn-primary">+ Create New Post</Link>
      </div>

      {posts.length === 0 ? (
        <p className="empty-message">No posts yet. Create your first post!</p>
      ) : (
        <table className="posts-table">
          <thead>
            <tr>
              <th>Title</th>
              <th>Status</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {posts.map((post) => (
              <tr key={post.id}>
                <td>{post.title}</td>
                <td>
                  <span className={`status-badge status-${post.status}`}>
                    {post.status}
                  </span>
                </td>
                <td>{new Date(post.created_at).toLocaleDateString()}</td>
                <td className="actions-cell">
                  <Link to={`/posts/${post.id}`} className="btn-link">View</Link>
                  <Link to={`/admin/posts/${post.id}/edit`} className="btn-link">Edit</Link>
                  <button onClick={() => handleDelete(post.id)} className="btn-link btn-danger">
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}