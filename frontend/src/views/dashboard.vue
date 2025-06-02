<template>
  <div class="dashboard">
    <h2>Create a Post</h2>
    <div class="card p-fluid">
      <div class="field">
        <label for="title">Title</label>
        <InputText id="title" v-model="newPost.title" class="w-full mb-2" />
      </div>
      <div class="field">
        <label for="content">Content</label>
        <Textarea id="content" v-model="newPost.content" rows="4" class="w-full mb-2" />
      </div>
      <Button label="Post" @click="createPost" />
    </div>

    <h2>Your Posts</h2>
    <div v-if="loading" class="text-center">Loading posts...</div>
    <div v-else>
      <div v-for="post in posts" :key="post.id" class="post-card mb-4">
        <div v-if="editingPostId !== post.id">
          <h3>{{ post.title }}</h3>
          <p class="post-content">{{ post.content }}</p>
          <div class="post-actions">
            <Button icon="pi pi-pencil" class="p-button-text p-button-sm" @click="startEditing(post)" />
            <Button icon="pi pi-trash" class="p-button-text p-button-sm p-button-danger" @click="confirmDelete(post.id)" />
          </div>
        </div>
        <div v-else>
          <InputText v-model="editingPost.title" class="w-full mb-2" />
          <Textarea v-model="editingPost.content" rows="4" class="w-full mb-2" />
          <div class="flex gap-2">
            <Button label="Save" @click="updatePost" />
            <Button label="Cancel" class="p-button-secondary" @click="cancelEditing" />
          </div>
        </div>
      </div>
    </div>

    <Toast />
    <ConfirmDialog />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import InputText from 'primevue/inputtext'
import Toast from 'primevue/toast'
import ConfirmDialog from 'primevue/confirmdialog'
import { useToast } from 'primevue/usetoast'
import { useConfirm } from 'primevue/useconfirm'

const toast = useToast()
const confirm = useConfirm()
const token = localStorage.getItem('token')

// Post data
const newPost = ref({
  title: '',
  content: ''
})
const posts = ref([])
const loading = ref(true)

// Editing state
const editingPostId = ref(null)
const editingPost = ref({
  title: '',
  content: ''
})

// Fetch all posts
const fetchPosts = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/posts', {
      headers: { Authorization: `Bearer ${token}` },
      withCredentials:true
    })
    posts.value = response.data
  } catch (err) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to fetch posts' })
  } finally {
    loading.value = false
  }
}

// Create new post
const createPost = async () => {
  if (!newPost.value.title.trim() || !newPost.value.content.trim()) {
    toast.add({ severity: 'warn', summary: 'Warning', detail: 'Title and content are required' })
    return
  }

  try {
    await axios.post('http://localhost:8080/api/posts/', newPost.value, {
      headers: { Authorization: `Bearer ${token}` }
    })
    toast.add({ severity: 'success', summary: 'Success', detail: 'Post created' })
    newPost.value = { title: '', content: '' }
    await fetchPosts()
  } catch (err) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to create post' })
  }
}

// Start editing post
const startEditing = (post) => {
  editingPostId.value = post.id
  editingPost.value = { ...post }
}

// Cancel editing
const cancelEditing = () => {
  editingPostId.value = null
  editingPost.value = { title: '', content: '' }
}

// Update post
const updatePost = async () => {
  try {
    await axios.put(`http://localhost:8080/api/posts/${editingPostId.value}`, editingPost.value, {
      headers: { Authorization: `Bearer ${token}` }
    })
    toast.add({ severity: 'success', summary: 'Success', detail: 'Post updated' })
    cancelEditing()
    await fetchPosts()
  } catch (err) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to update post' })
  }
}

// Delete post confirmation
const confirmDelete = (postId) => {
  confirm.require({
    message: 'Are you sure you want to delete this post?',
    header: 'Confirmation',
    icon: 'pi pi-exclamation-triangle',
    accept: () => deletePost(postId)
  })
}

// Delete post
const deletePost = async (postId) => {
  try {
    await axios.delete(`http://localhost:8080/api/posts/${postId}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    toast.add({ severity: 'success', summary: 'Success', detail: 'Post deleted' })
    await fetchPosts()
  } catch (err) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to delete post' })
  }
}

// Initialize
onMounted(fetchPosts)
</script>

<style scoped>
.dashboard {
  max-width: 800px;
  margin: 2rem auto;
  padding: 1rem;
}

.post-card {
  background: var(--surface-card);
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.post-content {
  white-space: pre-wrap;
  margin: 1rem 0;
}

.post-actions {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}

.field {
  margin-bottom: 1rem;
}
</style>