<script>
  import { onMount } from 'svelte';
  import { api, getUsuario } from './api.js';
  import Dashboard from './Dashboard.svelte';
  import Clientes from './Clientes.svelte';
  import Grupos from './Grupos.svelte';
  import Prestamos from './Prestamos.svelte';
  import Pagos from './Pagos.svelte';
  import Usuarios from './Usuarios.svelte';
  import Logs from './Logs.svelte';
  import { logout } from './api.js';

  let currentView = 'dashboard';
  let usuario = getUsuario();

  // Safety check: if usuario is null, logout
  if (!usuario) {
    logout();
    window.location.reload();
  }

  function handleLogout() {
    logout();
    window.location.reload();
  }

  function navigate(view) {
    currentView = view;
  }
</script>

<div class="app-layout">
  <aside class="sidebar">
    <div class="sidebar-header">
      <h1>MIAFIN</h1>
      <div class="user-info">
        <div>{usuario.nombre}</div>
        <div class="user-role">{usuario.rol}</div>
      </div>
    </div>

    <nav>
      <button 
        class:active={currentView === 'dashboard'} 
        on:click={() => navigate('dashboard')}
      >
        Dashboard
      </button>

      <button 
        class:active={currentView === 'clientes'} 
        on:click={() => navigate('clientes')}
      >
        Clientes
      </button>

      <button 
        class:active={currentView === 'grupos'} 
        on:click={() => navigate('grupos')}
      >
        Grupos
      </button>

      <button 
        class:active={currentView === 'prestamos'} 
        on:click={() => navigate('prestamos')}
      >
        Prestamos
      </button>

      <button 
        class:active={currentView === 'pagos'} 
        on:click={() => navigate('pagos')}
      >
        Pagos
      </button>

      {#if usuario.rol === 'admin'}
        <button 
          class:active={currentView === 'usuarios'} 
          on:click={() => navigate('usuarios')}
        >
          Usuarios
        </button>

        <button 
          class:active={currentView === 'logs'} 
          on:click={() => navigate('logs')}
        >
          Logs
        </button>
      {/if}
    </nav>

    <button class="logout-btn" on:click={handleLogout}>
      Cerrar Sesion
    </button>
  </aside>

  <main class="main-content">
    {#if currentView === 'dashboard'}
      <Dashboard />
    {:else if currentView === 'clientes'}
      <Clientes />
    {:else if currentView === 'grupos'}
      <Grupos />
    {:else if currentView === 'prestamos'}
      <Prestamos />
    {:else if currentView === 'pagos'}
      <Pagos />
    {:else if currentView === 'usuarios'}
      <Usuarios />
    {:else if currentView === 'logs'}
      <Logs />
    {/if}
  </main>
</div>

<style>
  .app-layout {
    display: flex;
    height: 100vh;
  }

  .sidebar {
    width: 250px;
    background: #2c3e50;
    color: white;
    display: flex;
    flex-direction: column;
  }

  .sidebar-header {
    padding: 20px;
    border-bottom: 1px solid #34495e;
  }

  h1 {
    margin: 0 0 15px 0;
    font-size: 24px;
  }

  .user-info {
    font-size: 14px;
  }

  .user-role {
    font-size: 12px;
    color: #95a5a6;
    text-transform: uppercase;
  }

  nav {
    flex: 1;
    padding: 20px 0;
  }

  nav button {
    width: 100%;
    padding: 12px 20px;
    background: transparent;
    color: white;
    border: none;
    text-align: left;
    cursor: pointer;
    font-family: 'Courier New', monospace;
  }

  nav button:hover {
    background: #34495e;
  }

  nav button.active {
    background: #3498db;
  }

  .logout-btn {
    margin: 20px;
    padding: 12px;
    background: #e74c3c;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-family: 'Courier New', monospace;
  }

  .logout-btn:hover {
    background: #c0392b;
  }

  .main-content {
    flex: 1;
    overflow-y: auto;
    background: #ecf0f1;
  }
</style>
