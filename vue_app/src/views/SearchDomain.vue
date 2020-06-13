<template id="">
  <div class="search_domain">
    <div class="main_container">
      <div class="search">
        <div class="form-group">
          <input class="form-input" v-model="domain" type="text" placeholder="domain.com">
        </div>
        <button v-if="validDomain" class="btn btn-primary" @click="search_domain">Search</button>
      </div>
      <div class="loading loading-lg showable" :class="{visible : isLoading}">
      </div>
      <div class="message showable" :class="{visible : failedStatus}">
        {{message}}
      </div>
      <div class="showable" :class="{visible : visiblePanel}">
        <div class="subpanel">
          <h2>Information</h2>
          <div class="card">
            <div class="line">
              <div class="left title">Servers have changed</div>
              <div class="right text_message" v-if="info.First">
                Still not available.
              </div>
              <div v-else>
                <div class="right">
                  {{info.Servers_changed}}
                </div>
              </div>
            </div>
            <div class="line">
              <div class="left title">SSL grade</div>
              <div class="right" v-if="info.SSL_grade">
                {{info.SSL_grade}}
              </div>
              <div v-else>
                <div class="right text_message">
                  Not provided.
                </div>
              </div>
            </div>
            <div class="line">
              <div class="left title">Previous SSL grade</div>
              <div class="right" v-if="info.Previous_SSL_grade">
                {{info.Previous_SSL_grade}}
              </div>
              <div v-else>
                <div class="right text_message">
                  Still not available.
                </div>
              </div>
            </div>
            <div class="line">
              <div class="left title">Is down</div>
              <div class="right">
                {{info.Is_down}}
              </div>
            </div>
            <div class="logo">
              <div class="title"><p>Logo</p></div>
              <div v-if="info.Logo">
                <div>

                </div>
                <img :src="info.Logo" alt="Logo is not available">

              </div>
              <div v-else>
                <div class="message">
                  Logo is not available.
                </div>
              </div>
            </div>
            <div class="block_line">
              <div class="title">Title</div>
              <div v-if="info.Title">
                {{info.Title}}
              </div>
              <div v-else>
                <div class="message">
                  Title is not available.
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="subpanel">
          <h2>Servers</h2>
          <div class="message showable" :class="{visible : noServers}">
            There are not related servers.
          </div>
          <div class="servers_list">
            <div
            v-for="server in info.Servers"
            :key="server.Address"
            >
            <div class="card">
              <div class="title">
                <p>Address</p>
              </div>
              <div class="content">
                <p>{{server.Address}}</p>
              </div>
              <div class="title">
                <p>Country</p>
              </div>
              <div class="content">
                <p>{{server.Country}}</p>
              </div>
              <div class="title">
                <p>Owner</p>
              </div>
              <div class="content">
                <p>{{server.Owner}}</p>
              </div>
              <div class="title">
                <p>SSL grade</p>
              </div>
              <div class="content">
                <div v-if="server.SSLGrade">
                  <p>{{server.SSLGrade}}</p>
                </div>
                <div v-else>
                  <div class="text_message">
                    SSL grade information is not available.
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="text/javascript">
import axios from "axios"
export default {
  name: 'Search_domain',
  data(){
    return {
      domain: "",
      info: "",

      message: "El ayuwoki es rial",

      isLoading: false,
      visiblePanel: false,
      failedStatus: false,
      noServers: false,
    }
  },
  computed:{
    validDomain(){
      let re = /^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$/g;
      if(re.exec(this.domain) != null ){
        return true;
      }
      else {
        return false;
      }
    }
  },
  methods:{
    search_domain(){
      this.noServers = false;
      this.isLoading = true;
      this.visiblePanel = false;
      this.failedStatus = false;
      var thisa = this;

      if(!this.validDomain){
        return
      }

      axios.get(process.env.VUE_APP_GOSERVER+'/querydomain?domain='+this.domain, { crossdomain: true })
      .then(function (response) {
        thisa.info = response.data;
        thisa.isLoading = false;
        // checking if there are no errors.
        if (thisa.info.Status.length == 0){
          thisa.visiblePanel = true;
          if (thisa.info.Servers.length == 0){
            thisa.noServers = true;
          }
        }
        else{
          // Showing status message
          thisa.failedStatus = true;
          thisa.message = thisa.info.Status;
        }
      })
      .catch(function (error) {
        // Showing error message
        thisa.failedStatus = true;
        thisa.isLoading = false;
        thisa.message = "Check your connection or try again later."
      })
    }
  }

}
</script>
<style media="screen">
.search_domain .main_container{
  margin-bottom: 200px;
}
.search_domain input{
  width: 100%;
}
.search_domain button{
  margin-top: 20px;
  width: 100%;
  font-size: 18px;
}
.search_domain .panel_info .panel-heading{
  font-size: 16px;
}
.search_domain .panel_info .data{
  display: block;
}
.search_domain .card{
  padding: 20px;
}
.search_domain .subpanel .line{
  position: relative;
  margin: 5px;
  height: 1.2em;
}
.search_domain .subpanel .line .left{
  left: 0;
  font-size: 1em;
  position:absolute;
}
.search_domain .subpanel .line .right{
  font-size: 0.9em;
  margin-top: 0.8em;
  position:absolute;
  right: 0;
}
.search_domain .subpanel{
  margin: 10px;
  margin-top: 40px;
}
.search_domain h2 {

}
.search_domain .showable {
  display: none;
}
.search_domain .showable.visible {
  display: block;
}
.search_domain .loading{
  margin: 20px;
  margin-top: 40px;
}
.search_domain .loading *{
  height: 200px;
}
.search_domain .message{
  color: #777;
  background: #eaeaea;
  border-radius: 10px;
  padding: 10px;
  margin-top: 10px;
}
.search_domain .text_message{
  color: #999;
  font-style: oblique;
  border-radius: 10px;
}
.search_domain .card .title{
  padding: 0.2em;
  font-size: 1em;
  font-weight: bold;
  height: 1.6em;
  margin: 0px;
  margin-top: 0.6em;
}
.search_domain .card .content{
  padding: 0.2em;
  margin: 0px;
  margin-top: 0.6em;
  height: 1.5em;
}
</style>
