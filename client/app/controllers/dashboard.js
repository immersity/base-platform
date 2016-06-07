// app/controllers/dashboard.js
import Ember from 'ember';

export default Ember.Controller.extend({
    session: Ember.inject.service('session'),
    actions: {
        printCurrentUser() {
            console.log(this.get('session').get('data'));
        },
        invalidateSession() {
            this.get('session').invalidate();
        }
    }
});
