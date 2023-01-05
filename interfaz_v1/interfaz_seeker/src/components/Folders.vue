<script>
export default {
  name: "Folders",
  props: {
    dataTemplate: Object,
  },
  methods: {
    showFolders() {
      let values = Object.keys(this.dataTemplate)
      return values
    }
  },
};
</script>

<template>
  <div class="px-10 py-10 w-auto h-screen">
    <div class="bg-black rounded-md shadow-lg shadow-lime-600 h-3/4 overflow-y-auto">
      <table class="table-auto text-white">
        <thead>
          <tr>
            <th>Folder1</th>
            <th>Folder 2</th>
            <th>File</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="folder in showFolders()">
            <td>
              <div class="flex flex-col items-center">
                <div class="flex items-center">
                  <img src="./images/carpeta1.png" alt="folder1" class="w-10 h-10 ">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
                  </svg>
                </div>
                <div class="text-sm">{{ folder.split(".")[0] }}</div>
              </div>
            </td>
            <td>
              <div class="flex flex-col items-center">
                <div class="flex items-center">
                  <img src="./images/carpeta2.png" alt="folder1" class="w-10 h-10 ">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
                  </svg>
                </div>
                <div class="text-sm">{{ folder.split(".")[1] }}</div>
              </div>
            </td>
            <td>
              <div class="flex flex-col items-center">
                <div class="w-10 h-10">
                  <img src="./images/documento.png" alt="folder1">
                </div>
                <div class="text-sm">{{ folder.split(".")[2] }}</div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

  </div>
</template>

