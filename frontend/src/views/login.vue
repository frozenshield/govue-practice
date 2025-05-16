<template>
  <div class="main-container">
    <Toast />
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
          <Button @click="handleLogin" label="Microsoft" icon="pi pi-microsoft" class="p-button-raised" />
          <Button @click="manualLogin" label="Manual Login" icon="pi pi-user" class="p-button-raised" />

          <!-- dialog -->
          <Dialog v-model:visible="isDialogLoginVisible" modal header="Login" :style="{ width: '30rem' }">
            <div class="dialog-input">
              <label for="username" class="font-semibold w-6rem">Username</label>
              <InputText id="username" v-model="loginForm.username" class="flex-auto" autocomplete="off" />
            </div>
            <div class="dialog-input">
              <label for="Password" class="font-semibold w-6rem">Password</label>
              <InputText id="password" v-model="loginForm.password" type="password" class="flex-auto" autocomplete="off" />
            </div>
            <div class="dialog-actions">
              <Button @click="handleManualLogin" label="login" class="p-button-raised" />
            </div>
          </Dialog>
        </div>
      </template>
    </Card>
  </div>
</template>


<script setup>
import { ref } from 'vue'
import { msalInstance, loginRequest } from '../msalConfig' // adjust path if needed
import {useRouter} from 'vue-router'
import {useToast} from 'primevue/usetoast'

import axios from 'axios'
import Card from 'primevue/card'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import Toast from 'primevue/toast';


const toast = useToast()
const router = useRouter()
const isDialogLoginVisible = ref(false)

const handleLogin = () => {
  msalInstance.loginRedirect(loginRequest)
}

const manualLogin = () => {
  isDialogLoginVisible.value = true
}

const loginForm = ref({
  username: '',
  password: ''
})

const handleManualLogin = async () => {
  try {
    const response = await axios.post('http://localhost:8080/api/users/login', {
      username: loginForm.value.username,
      password: loginForm.value.password
    })

    console.log('Login successful', response.data)

    const token = response.data.token
    localStorage.setItem('token', token)


    isDialogLoginVisible.value = false
    
    toast.add({
        severity: 'success',
        summary: 'Login Success',
        detail: 'You have logged in successfully',
        life: 3000
    })

    setTimeout(() => {
        router.push('/dashboard')
    }, 1000)
    
  } catch (error) {
    console.error('Login failed:', error)
   
    toast.add({
        severity: 'error',
        summary: 'Login Success',
        detail: 'Invalid username or password.',
        life: 3000
    })
  }
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

.dialog-input {
  display: flex;
  align-items: center;
  gap: 3rem;
  margin-bottom: 3rem;
}

.dialog-actions {
  display: flex;
  justify-content: center;
  gap: 1.5rem; /* adjust spacing as needed */
  margin-top: 1rem;
}
</style>