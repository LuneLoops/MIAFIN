<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let prestamos = [];
  let clientes = [];
  let grupos = [];
  let loading = true;
  let showForm = false;
  let form = { tipo: 'individual', cliente_id: null, grupo_id: null, monto: 0, tasa_interes: 0, plazo_meses: 0 };

  onMount(async () => {
    await loadData();
  });

  async function loadData() {
    loading = true;
    try {
      [prestamos, clientes, grupos] = await Promise.all([
        api.getPrestamos(),
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
      await api.createPrestamo(form);
      form = { tipo: 'individual', cliente_id: null, grupo_id: null, monto: 0, tasa_interes: 0, plazo_meses: 0 };
      showForm = false;
      await loadData();
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }

  async function descargarReporte(id) {
    try {
      const blob = await api.getReportePrestamoPDF(id);
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `prestamo-${id}.pdf`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }
</script>

<div class="page">
  <div class="page-header">
    <h2>Prestamos</h2>
    <button on:click={() => showForm = !showForm}>{showForm ? 'Cancelar' : 'Nuevo Prestamo'}</button>
  </div>

  {#if showForm}
    <div class="form-card">
      <h3>Nuevo Prestamo</h3>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
          <label>Tipo</label>
          <select bind:value={form.tipo} required>
            <option value="individual">Individual</option>
            <option value="comunal">Comunal</option>
          </select>
        </div>
        {#if form.tipo === 'individual'}
          <div class="form-group">
            <label>Cliente</label>
            <select bind:value={form.cliente_id} required>
              <option value={null}>Seleccionar...</option>
              {#each clientes as cliente}
                <option value={cliente.id}>{cliente.nombre}</option>
              {/each}
            </select>
          </div>
        {:else}
          <div class="form-group">
            <label>Grupo</label>
            <select bind:value={form.grupo_id} required>
              <option value={null}>Seleccionar...</option>
              {#each grupos as grupo}
                <option value={grupo.id}>{grupo.nombre}</option>
              {/each}
            </select>
          </div>
        {/if}
        <div class="form-group">
          <label>Monto (Bs.)</label>
          <input type="number" bind:value={form.monto} step="0.01" required />
        </div>
        <div class="form-group">
          <label>Tasa de Interes (%)</label>
          <input type="number" bind:value={form.tasa_interes} step="0.01" required />
        </div>
        <div class="form-group">
          <label>Plazo (meses)</label>
          <input type="number" bind:value={form.plazo_meses} required />
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
            <th>Tipo</th>
            <th>Cliente/Grupo</th>
            <th>Monto</th>
            <th>Tasa</th>
            <th>Plazo</th>
            <th>Saldo</th>
            <th>Estado</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          {#each prestamos as prestamo}
            <tr>
              <td>{prestamo.id}</td>
              <td>{prestamo.tipo}</td>
              <td>{prestamo.cliente?.nombre || prestamo.grupo?.nombre || '-'}</td>
              <td>Bs. {prestamo.monto.toFixed(2)}</td>
              <td>{prestamo.tasa_interes}%</td>
              <td>{prestamo.plazo_meses}m</td>
              <td>Bs. {prestamo.saldo_actual.toFixed(2)}</td>
              <td>{prestamo.estado}</td>
              <td>
                <button class="small-btn" on:click={() => descargarReporte(prestamo.id)}>PDF</button>
              </td>
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
  .small-btn { padding: 5px 10px; font-size: 12px; background: #28a745; }
  .small-btn:hover { background: #218838; }
  .form-card { background: white; padding: 20px; border-radius: 4px; border: 1px solid #ddd; margin-bottom: 20px; }
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
