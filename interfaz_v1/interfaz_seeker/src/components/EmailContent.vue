<script>
export default {
  name: "EmailConten",
  props: {
    dataTemplate: Object,
    messageEmailContent: String,
  },
  methods: {
    showEmail() {
      let values = Object.values(this.dataTemplate)
      return values[0]
    }
  },
};

</script>


<template>
  <div class="pt-10 pr-10 pb-28">
    <div class=" text-white h-screen">
      <div class="h-3/4 overflow-auto bg-black rounded-md shadow-lg shadow-lime-600">
        <div class="flex justify-center">
          <div class="pt-4 pb-4">
            <div v-if="messageEmailContent" class="mt-3 alert alert-secondary" v-html="dataTemplate[messageEmailContent]"></div>
            <span v-else v-html="showEmail()"></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

