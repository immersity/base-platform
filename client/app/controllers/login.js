import Ember from 'ember';

export default Ember.Controller.extend({
    session: Ember.inject.service('session'),
    actions: {
        authenticate: function () {
            const credentials = this.getProperties('identification', 'password');
            const authenticator = 'authenticator:token';

            this.get('session')
                .authenticate(authenticator, credentials)
                .catch((err) => {
                    this.set('errorMessage', err.error);
                });
        }
    }
});
