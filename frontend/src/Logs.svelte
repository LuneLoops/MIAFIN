<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let logs = [];
  let loading = true;

  onMount(async () => {
    loading = true;
    try {
      logs = await api.getLogs();
    } catch (err) {
      alert('Error: ' + err.message);
    } finally {
      loading = false;
    }
  });
</script>

<div class="page">
  <h2>Logs de Auditoria</h2>

  {#if loading}
    <p>Cargando...</p>
  {:else}
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Usuario</th>
            <th>Accion</th>
            <th>Entidad</th>
            <th>Descripcion</th>
            <th>Fecha</th>
          </tr>
        </thead>
        <tbody>
          {#each logs as log}
            <tr>
              <td>{log.id}</td>
              <td>{log.usuario?.nombre || '-'}</td>
              <td>{log.accion}</td>
              <td>{log.entidad}</td>
              <td>{log.descripcion}</td>
              <td>{new Date(log.created_at).toLocaleString()}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<style>
  .page { padding: 20px; }
  h2 { margin: 0 0 20px 0; }
  .table-container { background: white; border-radius: 4px; border: 1px solid #ddd; overflow-x: auto; }
  table { width: 100%; border-collapse: collapse; }
  th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; font-size: 14px; }
  th { background: #f8f9fa; font-weight: bold; }
  tbody tr:hover { background: #f8f9fa; }
</style>
