import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/CommonStyles.css';
import '../styles/SignUpPage.css';

const SignUpPage = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      alert("Passwords don't match!");
      return;
    }
    console.log({ name, email, password, confirmPassword });
    alert('Sign up successful!');
    navigate('/login');
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

          <button type="submit" className="button-base signup-button">
            Sign Up
          </button>
        </form>
      </div>
    </div>
  );
};

export default SignUpPage;
