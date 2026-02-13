<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let usuarios = [];
  let loading = true;
  let showForm = false;
  let form = { nombre: '', email: '', password: '', rol: 'asesor' };

  onMount(async () => {
    await loadData();
  });

  async function loadData() {
    loading = true;
    try {
      usuarios = await api.getUsuarios();
    } catch (err) {
      alert('Error: ' + err.message);
    } finally {
      loading = false;
    }
  }

  async function handleSubmit() {
    try {
      await api.createUsuario(form);
      form = { nombre: '', email: '', password: '', rol: 'asesor' };
      showForm = false;
      await loadData();
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }
</script>

<div class="page">
  <div class="page-header">
    <h2>Usuarios</h2>
    <button on:click={() => showForm = !showForm}>{showForm ? 'Cancelar' : 'Nuevo Usuario'}</button>
  </div>

  {#if showForm}
    <div class="form-card">
      <h3>Nuevo Usuario</h3>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
          <label>Nombre</label>
          <input bind:value={form.nombre} required />
        </div>
        <div class="form-group">
          <label>Email</label>
          <input type="email" bind:value={form.email} required />
        </div>
        <div class="form-group">
          <label>Contraseña</label>
          <input type="password" bind:value={form.password} required minlength="6" />
        </div>
        <div class="form-group">
          <label>Rol</label>
          <select bind:value={form.rol} required>
            <option value="asesor">Asesor</option>
            <option value="admin">Admin</option>
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
            <th>Email</th>
            <th>Rol</th>
          </tr>
        </thead>
        <tbody>
          {#each usuarios as usuario}
            <tr>
              <td>{usuario.id}</td>
              <td>{usuario.nombre}</td>
              <td>{usuario.email}</td>
              <td>{usuario.rol}</td>
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
  .form-card { background: white; padding: 20px; border-radius: 4px; border: 1px solid #ddd; margin-bottom: 20px; max-width: 600px; }
  h3 { margin: 0 0 20px 0; }
  .form-group { margin-bottom: 15px; }
  label { display: block; margin-bottom: 5px; font-weight: bold; }
  input, select { width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; font-family: 'Courier New', monospace; }
  .table-container { background: white; border-radius: 4px; border: 1px solid #ddd; overflow-x: auto; }
  table { width: 100%; border-collapse: collapse; }
  th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
  th { background: #f8f9fa; font-weight: bold; }
  tbody tr:hover { background: #f8f9fa; }
</style>
