<template>
  <div class="main-container">
    <div class="welcome-container">
      <h1 class="welcome-title">Social Media App</h1>
    </div>

    <Card class="centered-card">
      <template #header>
        <div class="card-header">
          <h2>Login</h2>
        </div>
      </template>
      <template #content>
        <div class="card-content">
          <p>Your premium experience starts here!</p>
          <Button @click="handleLogin" label="Microsoft" icon="pi pi-microsoft" class="p-button-raised" />
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup>
import { msalInstance, loginRequest } from './msalConfig' // adjust path if needed
import Card from 'primevue/card'
import Button from 'primevue/button'

const handleLogin = () => {
  msalInstance.loginRedirect(loginRequest)
}
</script>

<style scoped>
.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 2rem;
}

.welcome-container {
  margin-bottom: 2rem;
  text-align: center;
}

.welcome-title {
  color: #2c3e50;
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
  text-shadow: 1px 1px 2px rgba(0,0,0,0.1);
}

.centered-card {
  width: 100%;
  max-width: 500px;
  border-radius: 12px !important;
  box-shadow: 0 10px 20px rgba(0,0,0,0.1);
  overflow: hidden;
}

.card-header {
  background: linear-gradient(135deg, #6B73FF 0%, #000DFF 100%);
  color: white;
  padding: 1.5rem;
  text-align: center;
}

.card-header h2 {
  margin: 0;
  font-size: 1.8rem;
}

.card-content {
  padding: 2rem;
  text-align: center;
}

.card-content p {
  color: #555;
  font-size: 1.1rem;
  margin-bottom: 1.5rem;
}

/* PrimeVue button override */
.p-button {
  background: linear-gradient(135deg, #6B73FF 0%, #000DFF 100%);
  border: none;
}

.p-button:hover {
  background: linear-gradient(135deg, #5a62e0 0%, #000bc7 100%);
}
</style>