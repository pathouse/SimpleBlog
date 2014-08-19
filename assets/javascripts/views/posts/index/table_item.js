define([
  'jquery',
  'underscore',
  'backbone'
], function($, _, Backbone) {

  var PostsTableItemView = Backbone.View.extend({
    template: JST['templates/posts/index/table_item'],
    
    events: {
      'click .destroy': 'destroy'
    },
    
    tagName: 'tr',

    destroy: function() {
      this.model.destroy();
      this.remove();
      return false;
    },

    templateContext: function(){
      return {post: this.model}
    },

    render: function() {
      this.$el.html(this.template(this.templateContext()));
      return this;
    }
  });
  
  return PostsTableItemView;
});
  
