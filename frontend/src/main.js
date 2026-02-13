import './style.css';
import App from './App.svelte';
import Login from './Login.svelte';
import { getToken, getUsuario } from './api.js';

const app = (getToken() && getUsuario()) 
  ? new App({ target: document.getElementById('app') })
  : new Login({ target: document.getElementById('app') });

export default app;
