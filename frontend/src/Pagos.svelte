<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let prestamos = [];
  let loading = true;
  let form = { prestamo_id: null, monto: 0 };

  onMount(async () => {
    await loadData();
  });

  async function loadData() {
    loading = true;
    try {
      prestamos = await api.getPrestamos();
      prestamos = prestamos.filter(p => p.estado !== 'pagado' && p.saldo_actual > 0);
    } catch (err) {
      alert('Error: ' + err.message);
    } finally {
      loading = false;
    }
  }

  async function handleSubmit() {
    try {
      await api.createPago(form);
      form = { prestamo_id: null, monto: 0 };
      await loadData();
      alert('Pago registrado exitosamente');
    } catch (err) {
      alert('Error: ' + err.message);
    }
  }

  $: selectedPrestamo = prestamos.find(p => p.id === form.prestamo_id);
</script>

<div class="page">
  <h2>Registrar Pago</h2>

  <div class="form-card">
    <form on:submit|preventDefault={handleSubmit}>
      <div class="form-group">
        <label>Prestamo</label>
        <select bind:value={form.prestamo_id} required>
          <option value={null}>Seleccionar prestamo...</option>
          {#each prestamos as prestamo}
            <option value={prestamo.id}>
              #{prestamo.id} - {prestamo.cliente?.nombre || prestamo.grupo?.nombre} - Saldo: Bs. {prestamo.saldo_actual.toFixed(2)}
            </option>
          {/each}
        </select>
      </div>

      {#if selectedPrestamo}
        <div class="info-box">
          <p><strong>Saldo actual:</strong> Bs. {selectedPrestamo.saldo_actual.toFixed(2)}</p>
        </div>
      {/if}

      <div class="form-group">
        <label>Monto a pagar (Bs.)</label>
        <input type="number" bind:value={form.monto} step="0.01" required />
      </div>

      <button type="submit">Registrar Pago</button>
    </form>
  </div>
</div>

<style>
  .page { padding: 20px; }
  h2 { margin: 0 0 20px 0; }
  .form-card { background: white; padding: 20px; border-radius: 4px; border: 1px solid #ddd; max-width: 600px; }
  .form-group { margin-bottom: 15px; }
  label { display: block; margin-bottom: 5px; font-weight: bold; }
  input, select { width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; font-family: 'Courier New', monospace; }
  button { padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; font-family: 'Courier New', monospace; }
  button:hover { background: #0056b3; }
  .info-box { padding: 10px; background: #e7f3ff; border: 1px solid #b3d9ff; border-radius: 4px; margin-bottom: 15px; }
  .info-box p { margin: 0; }
</style>
