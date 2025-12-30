import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import BlogList from './pages/BlogList';
import PostDetail from './pages/PostDetail';
import Login from './pages/Login';
import AuthGuard from './components/AuthGuard';
import Dashboard from './pages/admin/Dashboard';
import AdminCreatePost from './pages/admin/CreatePost';
import EditPost from './pages/admin/EditPost';
import './App.css';

export default function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <Routes>
          {/* Public routes */}
          <Route path="/" element={<BlogList />} />
          <Route path="/posts/:id" element={<PostDetail />} />
          <Route path="/login" element={<Login />} />

          {/* Admin routes (protected) */}
          <Route path="/admin" element={<AuthGuard><Dashboard /></AuthGuard>} />
          <Route path="/admin/posts/new" element={<AuthGuard><AdminCreatePost /></AuthGuard>} />
          <Route path="/admin/posts/:id/edit" element={<AuthGuard><EditPost /></AuthGuard>} />

          {/* Fallback */}
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}