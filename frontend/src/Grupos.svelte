<script>
  import { onMount } from 'svelte';
  import { api, getUsuario } from './api.js';

  let grupos = [];
  let loading = true;
  let showForm = false;
  let form = { nombre: '', descripcion: '' };
  let usuario = getUsuario();

  onMount(async () => {
    await loadData();
  });

  async function loadData() {
    loading = true;
    try {
      grupos = await api.getGrupos();
    } catch (err) {
      alert('Error: ' + err.message);
    } finally {
      loading = false;
    }
  }

  async function handleSubmit() {
    try {
      await api.createGrupo(form);
      form = { nombre: '', descripcion: '' };
      showForm = false;
      await loadData();
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }
</script>

<div class="page">
  <div class="page-header">
    <h2>Grupos</h2>
    {#if usuario.rol === 'admin'}
      <button on:click={() => showForm = !showForm}>
        {showForm ? 'Cancelar' : 'Nuevo Grupo'}
      </button>
    {/if}
  </div>

  {#if showForm}
    <div class="form-card">
      <h3>Nuevo Grupo</h3>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
          <label>Nombre</label>
          <input bind:value={form.nombre} required />
        </div>
        <div class="form-group">
          <label>Descripcion</label>
          <textarea bind:value={form.descripcion} rows="3"></textarea>
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
            <th>Descripcion</th>
          </tr>
        </thead>
        <tbody>
          {#each grupos as grupo}
            <tr>
              <td>{grupo.id}</td>
              <td>{grupo.nombre}</td>
              <td>{grupo.descripcion || '-'}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<style>
  .page { padding: 20px; }
  .page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
  h2 { margin: 0; }
  button { padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; font-family: 'Courier New', monospace; }
  button:hover { background: #0056b3; }
  .form-card { background: white; padding: 20px; border-radius: 4px; border: 1px solid #ddd; margin-bottom: 20px; }
  h3 { margin: 0 0 20px 0; }
  .form-group { margin-bottom: 15px; }
  label { display: block; margin-bottom: 5px; font-weight: bold; }
  input, textarea { width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; font-family: 'Courier New', monospace; }
  .table-container { background: white; border-radius: 4px; border: 1px solid #ddd; overflow-x: auto; }
  table { width: 100%; border-collapse: collapse; }
  th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
  th { background: #f8f9fa; font-weight: bold; }
  tbody tr:hover { background: #f8f9fa; }
</style>
