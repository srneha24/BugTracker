import axiosInstance from '../api/axiosInstance';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/CommonStyles.css';
import '../styles/SignUpPage.css';

const SignUpPage = () => {
  const [name, setName] = useState('');
  const [username, setUsername] = useState(''); 
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (password !== confirmPassword) {
      setError("Passwords don't match!");
      return;
    }

    try {
      const response = await axiosInstance.post('/user/signup', {
        name,
        username,
        email,
        password,
      });

      if (response.data.success) {
        if (response.data.token) {
          localStorage.setItem('token', response.data.token);
        }
        alert('Sign up successful!');
        navigate('/home');
      } else {
        setError(response.data.message || 'Sign up failed');
      }
    } catch (err: any) {
      if (err.response && err.response.data) {
        setError(err.response.data.message || 'Sign up failed');
      } else {
        setError('Network error or server is down');
      }
    }
  };

  return (
    <div className="signup-container">
      <div className="signup-card">
        <h2 className="signup-title">Sign Up</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label className="input-label">Name</label><br />
            <input
              type="text"
              value={name}
              onChange={e => setName(e.target.value)}
              required
              placeholder="Your full name"
              className="input-field"
            />
          </div>

          <div className="form-group">
            <label className="input-label">Username</label><br />
            <input
              type="text"
              value={username}
              onChange={e => setUsername(e.target.value)}
              required
              placeholder="Choose a username"
              className="input-field"
            />
          </div>

          <div className="form-group">
            <label className="input-label">Email</label><br />
            <input
              type="email"
              value={email}
              onChange={e => setEmail(e.target.value)}
              required
              placeholder="you@example.com"
              className="input-field"
            />
          </div>

          <div className="form-group">
            <label className="input-label">Password</label><br />
            <input
              type="password"
              value={password}
              onChange={e => setPassword(e.target.value)}
              required
              placeholder="Create a password"
              className="input-field"
            />
          </div>

          <div className="form-group">
            <label className="input-label">Confirm Password</label><br />
            <input
              type="password"
              value={confirmPassword}
              onChange={e => setConfirmPassword(e.target.value)}
              required
              placeholder="Confirm your password"
              className="input-field"
            />
          </div>

          {error && <p style={{ color: 'red' }}>{error}</p>}

          <button type="submit" className="button-base signup-button">
            Sign Up
          </button>
        </form>
      </div>
    </div>
  );
};

export default SignUpPage;
