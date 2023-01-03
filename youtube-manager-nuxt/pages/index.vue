<template>
  <section class="section">
    <div class="container">
      <div class="block video-block" v-for="item in items">
        <AppVideo :item="item" :videoId="item.id" />
      </div>
    </div>
  </section>
</template>

<script>
  import ROUTES from '~/routes/api'
  export default{
    computed: {
      items() {
        return this.$store.getters.getPopularVideos
      }
    },

    async fetch({store}) {
      const payload = {
        uri: ROUTES.GET.POPULARS
      }

      if (store.getters.getPopularVideos && store.getters.getPopularVideos.length > 0) {
        return
      }

      await store.dispatch('fetchPopularVideos', payload)
    }
  }
</script>

<style scoped>
  .video-block {
    max-width: 900px;
  }
</style>
