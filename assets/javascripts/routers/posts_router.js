define([
  'jquery',
  'underscore',
  'backbone',
  'collections/posts',
  'views/posts/index'//,
//  'views/posts/show',
//  'views/posts/new',
//  'views/posts/edit'
], function($, _, Backbone, PostsCollection, PostsIndexView) {//, PostsShowView, PostsNewView, PostsEditView){

  var PostsRouter = Backbone.Router.extend({
    
    initialize: function() {
      this.collection = new PostsCollection();
    },

    routes: {
      'posts': 'postsIndex',
      'posts/': 'postsIndex',
//      'posts/new': 'postsNew',
//      'posts/:id/edit': 'postsEdit'
//      'posts/:id': 'postsShow'
    },

    postsIndex: function() {
      var view = new PostsIndexView({
        collection: this.collection
      });
      $('#content').html(view.render().el);
    }//,

    // postsNew: function() {
    //   var view = new PostsNewView({
    //     collection: this.collection
    //   });
    //   $('#content').html(view.render().el);
    // },
    
    // postsEdit: function(id) {
    //   var post = new this.collection.model({id: id});
    //   var view  = new PostsEditView({
    //     model: post
    //   });
    //   $('#content').html(view.render().el);
    // },

    // postsShow: function(id) {
    //   var post = new this.collection.model({id: id});
    //   var view = new PostsShowView({
    //     model: post
    //   });
    //   $('#content').html(view.render().el);
    // }
  });
  
  return PostsRouter;
});
        
    
