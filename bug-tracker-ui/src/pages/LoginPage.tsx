import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/CommonStyles.css';
import '../styles/LoginPage.css';

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    console.log({ email, password });
    alert('Login successful!');
    navigate('/home');
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <h2 className="login-title">Bug Tracker</h2>

        <form onSubmit={handleLogin}>
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
              placeholder="Enter your password"
              className="input-field"
            />
          </div>

          <button type="submit" className="button-base login-button">
            Login
          </button>
        </form>

        <button onClick={() => navigate('/signup')} className="button-base signup-link-button">
          Sign Up
        </button>
      </div>
    </div>
  );
};

export default LoginPage;
