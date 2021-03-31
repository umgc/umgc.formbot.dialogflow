Vue.component('forms', {
  props: {
    list: Array,
    filterKey: String
  },
  data: function(){
    return{
      docView: false,
      docViererURLid:""
    }
  },
  computed: { 
    filteredForms: function () {
      var filterKey = this.filterKey && this.filterKey.toLowerCase();
      var data = this.list
      if (filterKey) {
        data = data.filter(function (row) {
          return Object.keys(row).some(function (key) {
            return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
          })
        })
      }
      return data;
    }
  },
  methods:{
    readForm(form){
      this.$emit('openformreader',form.id);
    },
    editForm(form){
      this.$emit('openformeditor',form.id);
    },
    copyTextArea(txt) {
        navigator.clipboard.writeText(txt);
    },
    formURL(id){
      return "https://docs.google.com/document/d/" + id;
    },
    toggelDocViewer(id){
      var self = this;
      if(self.docView){
        self.docView = false;
      }else{
        self.docView = true;
        self.docViererURLid = id;
      }
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  /* TODO
   * - Clean up layout
   * - Add toggel to show and hide the iframe
   */
  template: `  
  <article class="list" style="">
    <section class="view" v-if="!docView">
      <ul>
        <li v-for="form in filteredForms" class="cf doc">
          <div class="title">{{form.name}}</div>
          <div class="btn" @click="copyTextArea(toggelDocViewer(form.id))">View Doc</div> 
          <div class="btn" @click="copyTextArea(formURL(form.id))">URL</div>        
        </li>
      </ul>
    </section>
    <section class="doc view" v-if="docView">
      <div class="btn" @click="copyTextArea(toggelDocViewer(docViererURLid))">Back to List</div> 
      <iframe :src="formURL(docViererURLid)"></iframe>
    </section>
  </article>
      `
    })