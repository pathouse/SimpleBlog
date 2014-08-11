define([
  'jquery',
  'underscore', 
  'backbone',
  'views/posts/index/table_item'
], function($, _, Backbone, PostsIndexTableItemView){

  var PostsIndexView = Backbone.View.extend({
    template: JST["templates/posts/index"],
    
    intialize: function() {
      this.collection.on('sync', this.render, this);
      this.collection.fetch();
    },

    addAll: function() {
      this.collection.each(this.addOne, this);
    },
    
    addOne: function(post) {
      var view = new PostsIndexTableItemView({
        model: post
      });
      this.$('tbody').append(view.render().el);
    },

    render: function() {
      $(this.el).html(this.template({}));
      this.allAll();
      return this;
    }
  });

  return PostsIndexView;
});
