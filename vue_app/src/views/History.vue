<template id="">
  <div class="history">
    <div class="message showable" :class="{visible : !showTable}">
        {{message}}
    </div>
    <div class="formatted_table showable" :class="{visible : showTable}">
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
      history : [],
      failedStatus : false,
      showTable: false,
      message: "",
    }
  },
  computed:{
  },
  methods:{
    view_queries(){
      this.message = ""
      this.failedStatus = false
      var thisa = this;
      this.showTable = false
      axios.get(GOSERVER+'/viewed_domains')
      .then(function (response) {
        thisa.history = response.data;
        if (thisa.history.Queries != null ){
         thisa.showTable = true
        }
        else{
          thisa.message = "You haven't query a domain."
        }
      })
      .catch(function (error) {
        console.log(error)
        thisa.failedStatus = true
        thisa.message = "Error connecting with the server"
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
.history .showable {
  display: none;
}
.history .showable.visible {
  display: block;
}
.history .message{
  color: #777;
  background: #eaeaea;
  border-radius: 10px;
  padding: 10px;
  margin: 20px;
  margin-top: 30px;
}
</style>
