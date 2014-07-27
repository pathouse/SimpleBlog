$(function() {
    var User = Backbone.Model.extend({
        validate: function(attributes) {
            if (attributes.byline === undefined) {
                return "Byline required for user.";
            } else if (attributes.email === undefined) {
                return "Email required for user.";
            }
        },

        initialize: function() {
            console.log('New user initialized');
            
            this.on('invalid', function(model, error) {
                console.log(error);
            });
            
            this.on('change:byline', function() {
                console.log("Byline for user has been changed to " + this.get("byline"));
            });
            
            this.on('change:email', function() {
                console.log("Email for user has been changed to " + this.get("email"));
            });                   
        }
    });

    var Post = Backbone.Model.extend({
        validate: function(attributes) {
            if (attributes.title === undefined) {
                return "Title required for post.";
            } else if (attributes.body === undefined) {
                return "Body required for post.";
            } else if (attributes.userId === undefined) {
                return "User ID required for post.";
            }
        },

        initialize: function() {
            console.log('New post initialized');
            
            this.on('change:title', function() {
                console.log("Title for post has been changed to " + this.get("title"));
            });

            this.on('change:draft', function() {
                console.log("Draft status for post has been changed to " + this.get("draft"));
            });
        }
    });

    
});
