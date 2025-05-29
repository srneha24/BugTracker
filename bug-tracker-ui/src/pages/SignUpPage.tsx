import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

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
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f0f0f0',
      }}
    >
      <div
        style={{
          backgroundColor: '#ccc',
          padding: '3rem 2rem',
          borderRadius: '8px',
          width: '480px',
          boxShadow: '0 4px 8px rgba(0,0,0,0.1)',
        }}
      >
        <h2
          style={{
            fontSize: '3rem',
            fontWeight: '700',
            marginBottom: '1.5rem',
            textAlign: 'center',
          }}
        >
          Sign Up
        </h2>
        <form onSubmit={handleSubmit}>
          <div style={{ marginBottom: '1rem' }}>
            <label style={{ fontWeight: '600' }}>Name</label><br />
            <input
              type="text"
              value={name}
              onChange={e => setName(e.target.value)}
              required
              placeholder="Your full name"
              style={{
                width: '100%',
                padding: '0.5rem',
                backgroundColor: 'white',
                border: '1px solid #999',
                borderRadius: '4px',
                boxSizing: 'border-box',
              }}
            />
          </div>

          <div style={{ marginBottom: '1rem' }}>
            <label style={{ fontWeight: '600' }}>Email</label><br />
            <input
              type="email"
              value={email}
              onChange={e => setEmail(e.target.value)}
              required
              placeholder="you@example.com"
              style={{
                width: '100%',
                padding: '0.5rem',
                backgroundColor: 'white',
                border: '1px solid #999',
                borderRadius: '4px',
                boxSizing: 'border-box',
              }}
            />
          </div>

          <div style={{ marginBottom: '1rem' }}>
            <label style={{ fontWeight: '600' }}>Password</label><br />
            <input
              type="password"
              value={password}
              onChange={e => setPassword(e.target.value)}
              required
              placeholder="Create a password"
              style={{
                width: '100%',
                padding: '0.5rem',
                backgroundColor: 'white',
                border: '1px solid #999',
                borderRadius: '4px',
                boxSizing: 'border-box',
              }}
            />
          </div>

          <div style={{ marginBottom: '1rem' }}>
            <label style={{ fontWeight: '600' }}>Confirm Password</label><br />
            <input
              type="password"
              value={confirmPassword}
              onChange={e => setConfirmPassword(e.target.value)}
              required
              placeholder="Confirm your password"
              style={{
                width: '100%',
                padding: '0.5rem',
                backgroundColor: 'white',
                border: '1px solid #999',
                borderRadius: '4px',
                boxSizing: 'border-box',
              }}
            />
          </div>

          <button
            type="submit"
            style={{
              padding: '0.75rem',
              width: '100%',
              cursor: 'pointer',
              backgroundColor: '#000000',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              fontWeight: '600',
              fontSize: '1rem',
              marginTop: '2rem',
            }}
          >
            Sign Up
          </button>
        </form>
      </div>
    </div>
  );
};

export default SignUpPage;
