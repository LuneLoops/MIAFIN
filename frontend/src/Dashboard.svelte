<script>
  import { onMount } from 'svelte';
  import { api } from './api.js';

  let liquidez = null;
  let loading = true;
  let error = '';

  onMount(async () => {
    try {
      liquidez = await api.getLiquidez();
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  });

  async function descargarReporte() {
    try {
      const blob = await api.getReporteGeneralPDF();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'reporte-general.pdf';
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (err) {
      alert('Error descargando reporte: ' + err.message);
    }
  }
</script>

<div class="dashboard">
  <h2>Dashboard</h2>

  {#if loading}
    <p>Cargando...</p>
  {:else if error}
    <div class="error">{error}</div>
  {:else if liquidez}
    <div class="metrics">
      <div class="metric">
        <div class="metric-label">Total Recuperado</div>
        <div class="metric-value">Bs. {liquidez.total_recuperado.toFixed(2)}</div>
      </div>

      <div class="metric">
        <div class="metric-label">Total Saldo Pendiente</div>
        <div class="metric-value">Bs. {liquidez.total_saldo_pendiente.toFixed(2)}</div>
      </div>

      <div class="metric highlight">
        <div class="metric-label">Liquidez Actual</div>
        <div class="metric-value">Bs. {liquidez.liquidez.toFixed(2)}</div>
      </div>
    </div>

    <button on:click={descargarReporte} class="report-btn">
      Descargar Reporte General PDF
    </button>
  {/if}
</div>

<style>
  .dashboard {
    padding: 20px;
  }

  h2 {
    margin: 0 0 20px 0;
  }

  .metrics {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 20px;
    margin-bottom: 20px;
  }

  .metric {
    padding: 20px;
    background: white;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  .metric.highlight {
    background: #e7f3ff;
    border-color: #007bff;
  }

  .metric-label {
    font-size: 14px;
    color: #666;
    margin-bottom: 10px;
  }

  .metric-value {
    font-size: 24px;
    font-weight: bold;
  }

  .report-btn {
    padding: 12px 24px;
    background: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-family: 'Courier New', monospace;
  }

  .report-btn:hover {
    background: #218838;
  }

  .error {
    padding: 10px;
    background: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
    border-radius: 4px;
  }
</style>
