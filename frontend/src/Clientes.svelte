<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let clientes = [];
  let grupos = [];
  let loading = true;
  let showForm = false;
  let form = { nombre: '', ci: '', telefono: '', grupo_id: null };

  onMount(async () => {
    await loadData();
  });

  async function loadData() {
    loading = true;
    try {
      [clientes, grupos] = await Promise.all([
        api.getClientes(),
        api.getGrupos()
      ]);
    } catch (err) {
      alert('Error: ' + err.message);
    } finally {
      loading = false;
    }
  }

  async function handleSubmit() {
    try {
      await api.createCliente(form);
      form = { nombre: '', ci: '', telefono: '', grupo_id: null };
      showForm = false;
      await loadData();
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }
</script>

<div class="page">
  <div class="page-header">
    <h2>Clientes</h2>
    <button on:click={() => showForm = !showForm}>
      {showForm ? 'Cancelar' : 'Nuevo Cliente'}
    </button>
  </div>

  {#if showForm}
    <div class="form-card">
      <h3>Nuevo Cliente</h3>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
          <label>Nombre</label>
          <input bind:value={form.nombre} required />
        </div>
        <div class="form-group">
          <label>CI</label>
          <input bind:value={form.ci} required />
        </div>
        <div class="form-group">
          <label>Telefono</label>
          <input bind:value={form.telefono} />
        </div>
        <div class="form-group">
          <label>Grupo (opcional)</label>
          <select bind:value={form.grupo_id}>
            <option value={null}>Sin grupo</option>
            {#each grupos as grupo}
              <option value={grupo.id}>{grupo.nombre}</option>
            {/each}
          </select>
        </div>
        <button type="submit">Guardar</button>
      </form>
    </div>
  {/if}

  {#if loading}
    <p>Cargando...</p>
  {:else}
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
            <th>CI</th>
            <th>Telefono</th>
            <th>Grupo</th>
          </tr>
        </thead>
        <tbody>
          {#each clientes as cliente}
            <tr>
              <td>{cliente.id}</td>
              <td>{cliente.nombre}</td>
              <td>{cliente.ci}</td>
              <td>{cliente.telefono || '-'}</td>
              <td>{cliente.grupo?.nombre || '-'}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<style>
  .page {
    padding: 20px;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  h2 {
    margin: 0;
  }

  button {
    padding: 10px 20px;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-family: 'Courier New', monospace;
  }

  button:hover {
    background: #0056b3;
  }

  .form-card {
    background: white;
    padding: 20px;
    border-radius: 4px;
    border: 1px solid #ddd;
    margin-bottom: 20px;
  }

  h3 {
    margin: 0 0 20px 0;
  }

  .form-group {
    margin-bottom: 15px;
  }

  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }

  input, select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-family: 'Courier New', monospace;
  }

  .table-container {
    background: white;
    border-radius: 4px;
    border: 1px solid #ddd;
    overflow-x: auto;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  th, td {
    padding: 12px;
    text-align: left;
    border-bottom: 1px solid #ddd;
  }

  th {
    background: #f8f9fa;
    font-weight: bold;
  }

  tbody tr:hover {
    background: #f8f9fa;
  }
</style>
