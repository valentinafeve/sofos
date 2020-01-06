<template>
  <div id="app">
    
    <!-- Tabs -->
    <ul class="tab tab-block">
      <li class="tab-item" :class="{active : tab1_active}" @click="search">
        <a href="#">Search</a>
      </li>
      <li class="tab-item" :class="{active : tab2_active}" @click="history">
        <a href="#">History</a>
      </li>
    </ul>

    <!-- Views -->
    <router-view />

  </div>
</template>

<script>

export default {
  name: 'app',
  data(){
    return {
      tab1_active : false,
      tab2_active : false
    }
  },
  methods:{
    search(){
      this.tab1_active = true
      this.tab2_active = false
      this.$router.push('search')
    },
    history(){
      this.tab2_active = true
      this.tab1_active = false
      this.$router.push('history')
    }
  },
  components: {
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
