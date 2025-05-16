<template>
  <div class="dashboard">
    <h2>Create a Post</h2>

    <Textarea v-model="content" rows="4" class="w-full mb-2" placeholder="What's on your mind?" />
    <Button label="Post" @click="createPost" />

    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'

const toast = useToast()
const content = ref('')
const token = localStorage.getItem('token')


const createPost = async () => {
  if (!content.value.trim()) return

  try {
    await axios.post('http://localhost:8080/api/posts', { content: content.value }, {
      headers: { Authorization: `Bearer ${token}` }
    })
    toast.add({ severity: 'success', summary: 'Success', detail: 'Post created' })
    content.value = ''
  } catch (err) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to create post' })
  }
}
</script>


<style scoped>
.dashboard {
  max-width: 600px;
  margin: 2rem auto;
  padding: 1rem;
}
</style>
