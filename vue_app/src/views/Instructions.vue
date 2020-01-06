<template id="">
  <div class="instructions">
    <div class="main_container">
      <h3>
        About
      </h3>
      <div class="card">
        <h5>Search domain</h5>
        <div class="">
          <p>Type a domain without the www ot https, example:</p>
        </div>
      </div>
      <div class="card">
        <h5>History</h5>
        <div class="">
          <p>Queries will obly be saved if the query was succesul.</p>
          <p>The date is saved according to the server clock (+1)</p>
        </div>
      </div>

    </div>
  </div>
</template>
<script type="text/javascript">
import axios from "axios"
export default {
  name: "Instructions",
  data(){
    return {
      history : {}
    }
  },
  methods:{
  },
  mounted(){
  },
}
</script>
<style media="screen">
.instructions .card{
  padding: 20px;
  margin-bottom: 20px;
  margin-top: 20px;
  border-radius: 20px;
  -webkit-box-shadow: 0px 0px 6px -4px rgba(200,200,200,0.5);
  -moz-box-shadow: 0px 0px 6px -4px rgba(200,200,200,0.5);
  box-shadow: 0px 0px 6px -4px rgba(200,200,200,0.5);

}
</style>
