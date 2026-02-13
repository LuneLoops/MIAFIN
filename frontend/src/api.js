const API_URL = 'http://localhost:8080';

let token = localStorage.getItem('token') || '';
let usuario = JSON.parse(localStorage.getItem('usuario') || 'null');

export function setToken(newToken) {
  token = newToken;
  localStorage.setItem('token', newToken);
}

export function setUsuario(user) {
  usuario = user;
  localStorage.setItem('usuario', JSON.stringify(user));
}

export function getToken() {
  return token;
}

export function getUsuario() {
  return usuario;
}

export function logout() {
  token = '';
  usuario = null;
  localStorage.removeItem('token');
  localStorage.removeItem('usuario');
}

async function request(method, endpoint, body = null) {
  const options = {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
  };

  if (token) {
    options.headers['Authorization'] = `Bearer ${token}`;
  }

  if (body) {
    options.body = JSON.stringify(body);
  }

  const response = await fetch(`${API_URL}${endpoint}`, options);
  
  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: 'Error desconocido' }));
    throw new Error(error.error || 'Error en la solicitud');
  }

  // Para PDFs
  if (response.headers.get('Content-Type') === 'application/pdf') {
    return await response.blob();
  }

  return await response.json();
}

export const api = {
  login: (email, password) => request('POST', '/login', { email, password }),
  
  // Usuarios
  getUsuarios: () => request('GET', '/usuarios'),
  createUsuario: (data) => request('POST', '/usuarios', data),
  
  // Clientes
  getClientes: () => request('GET', '/clientes'),
  createCliente: (data) => request('POST', '/clientes', data),
  
  // Grupos
  getGrupos: () => request('GET', '/grupos'),
  createGrupo: (data) => request('POST', '/grupos', data),
  
  // Préstamos
  getPrestamos: () => request('GET', '/prestamos'),
  createPrestamo: (data) => request('POST', '/prestamos', data),
  
  // Pagos
  createPago: (data) => request('POST', '/pagos', data),
  
  // Reportes
  getLiquidez: () => request('GET', '/reportes/liquidez'),
  getReporteGeneralPDF: () => request('GET', '/reportes/general/pdf'),
  getReportePrestamoPDF: (id) => request('GET', `/reportes/prestamo/${id}/pdf`),
  getComprobantePagoPDF: (id) => request('GET', `/reportes/pago/${id}/pdf`),
  
  // Logs
  getLogs: () => request('GET', '/logs'),
};
