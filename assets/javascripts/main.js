require.config({
    paths: {
        jquery: 'vendor/jquery-2.1.1.min',
        underscore: 'vendor/underscore-min',
        backbone: 'vendor/backbone-min'
    }
});

require([
    'application',
], function(App){
    App.initialize();
});
        
