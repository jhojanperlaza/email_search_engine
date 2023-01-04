<script setup>
import Folders from './components/Folders.vue';
import EmailConten from './components/EmailContent.vue';

</script>

<template>
  <header>
    <!-- navbar component-->
    <nav class="bg-black py-2">
      <div class="max-w-7xl mx-auto">
        <div class="flex justify-between items-center">

          <!--logo-->
          <div>
            <img src="./components/images/logo.png" alt="logo" class="w-64 h-auto">
          </div>

          <!--primary nav search-->
          <div class="w-96">
            <form v-on:submit.prevent="sendToBackend">
              <label for="default-search"
                class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                  <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                  </svg>
                </div>
                <input type="search" id="default-search" v-model="query_search"
                  class="block w-full p-4 pl-10 text-sm text-white border border-gray-300 rounded-lg bg-gray-700 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="Search..." required>
                <button type="submit"
                  class="text-white absolute right-2.5 bottom-2.5 bg-lime-600 hover:bg-lime-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
              </div>
            </form>
          </div>

          <!--secondary nav icons-->
          <div class="hidden md:flex space-x-4">
            <a href="https://www.linkedin.com/in/jhojan-p-a45b1813a/">
              <img src="./components/icons/linkedin.png" alt="linkedin-icon" class="w-14 h-14">
            </a>
            <a href="https://github.com/jhojanperlaza">
              <img src="./components/icons/github.png" alt="linkedin-icon" class="w-14 h-14">
            </a>
          </div>

          <!--moblie button-->
          <div class="md:hidden flex pr-6" style="color:aliceblue">
            <button class="mobile-menu-button">
              <svg class="w-11 h-11" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
              </svg>
            </button>
          </div>

        </div>
      </div>


      <!--mobile menu-->
    </nav>
  </header>
  <main>
    <div class="flex justify-between">
      <div class="pt-10 pl-10 pr-2">
        <Folders :dataTemplate="responseToComponents"/>
      </div>
      <div class="w-8/12">
        <EmailConten :dataTemplate="responseToComponents"/>
      </div>
    </div>
  </main>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DataSearch',

  data() {
    return {
      query_search: '',
      responseToComponents: 'aun nada',
    }
  },

  methods: {
    sendToBackend() {
      axios.post("http://localhost:3000/api/searchQuery", this.query_search)
        .then((response) => {
          console.log(response.data)
          this.responseToComponents = response.data
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
    },
  },
  components: {
    EmailConten,
    Folders,
  },
}
</script>
