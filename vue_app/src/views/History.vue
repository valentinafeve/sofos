<template id="">
  <div class="">
    <div class="formatted_table">
      <table class="table">
        <thead>
          <tr>
            <th>Domain</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          <tr class="active" v-for="query in history.Queries" :key="query.Time">
            <td>{{query.Domain}}</td>
            <td>{{query.Time}}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script type="text/javascript">
import axios from "axios"
export default {
  name: "History",
  data(){
    return {
      history : {}
    }
  },
  methods:{
    view_queries(){
      var thisa = this;
      axios.get(GO_SERVER+'/viewed_domains')
      .then(function (response) {
        // Saving response
        thisa.history = response.data;
        console.log(thisa.history);
      })
      .catch(function (error) {
        console.log(error);
      })
    }
  },
  mounted(){
    // Send get request
    this.view_queries();
  },
}
</script>
<style media="screen">
.formatted_table{
  margin: 30px;
}
</style>
