<script setup>
import Navbar from './components/Navbar.vue'
import Folders from './components/Folders.vue';
import EmailConten from './components/EmailContent.vue';

</script>

<template>
  <header>
    <Navbar />
  </header>
  <main>
    <div class="flex justify-between">
      <div class="pt-10 pl-10 pr-2">
        <Folders />
      </div>
      <div class="w-8/12">
        <EmailConten />
      </div>
    </div>
  </main>
</template>
