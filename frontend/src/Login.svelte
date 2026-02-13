<script>
  import { api, setToken, setUsuario } from './api.js';

  let email = '';
  let password = '';
  let error = '';
  let loading = false;

  async function handleSubmit() {
    error = '';
    loading = true;
    
    try {
      const response = await api.login(email, password);
      setToken(response.token);
      setUsuario(response.usuario);
      window.location.reload();
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>

<div class="login-container">
  <div class="login-box">
    <h1>MIAFIN</h1>
    <p>Sistema Financiero Comunal</p>

    <form on:submit|preventDefault={handleSubmit}>
      <div class="form-group">
        <label for="email">Email</label>
        <input 
          id="email"
          type="email" 
          bind:value={email} 
          required 
          disabled={loading}
        />
      </div>

      <div class="form-group">
        <label for="password">Contraseña</label>
        <input 
          id="password"
          type="password" 
          bind:value={password} 
          required 
          disabled={loading}
        />
      </div>

      {#if error}
        <div class="error">{error}</div>
      {/if}

      <button type="submit" disabled={loading}>
        {loading ? 'Ingresando...' : 'Ingresar'}
      </button>
    </form>

    <div class="demo-info">
      <p>Usuario demo: admin@miafin.local</p>
      <p>Contraseña: admin123</p>
    </div>
  </div>
</div>

<style>
  .login-container {
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    background: #f5f5f5;
  }

  .login-box {
    background: white;
    padding: 40px;
    border-radius: 4px;
    border: 1px solid #ddd;
    width: 100%;
    max-width: 400px;
  }

  h1 {
    margin: 0 0 10px 0;
    font-size: 32px;
  }

  p {
    margin: 0 0 30px 0;
    color: #666;
  }

  .form-group {
    margin-bottom: 20px;
  }

  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }

  input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-family: 'Courier New', monospace;
  }

  input:disabled {
    background: #f5f5f5;
  }

  button {
    width: 100%;
    padding: 12px;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    font-family: 'Courier New', monospace;
    font-weight: bold;
    cursor: pointer;
  }

  button:hover:not(:disabled) {
    background: #0056b3;
  }

  button:disabled {
    background: #ccc;
    cursor: not-allowed;
  }

  .error {
    padding: 10px;
    background: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
    border-radius: 4px;
    margin-bottom: 20px;
  }

  .demo-info {
    margin-top: 20px;
    padding-top: 20px;
    border-top: 1px solid #ddd;
    font-size: 12px;
    color: #666;
  }

  .demo-info p {
    margin: 5px 0;
  }
</style>
