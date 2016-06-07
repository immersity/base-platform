// app/controllers/sign-up.js
import Ember from 'ember';

export default Ember.Controller.extend({
    email: '',
    password: '',
    firstName: '',
    lastName: '',
    actions: {
        signUp: function () {
            const candidate = this.store.createRecord('account', {
                email: this.get('email'),
                password: this.get('password'),
                firstName: this.get('firstName'),
                lastName: this.get('lastName'),
            });

            candidate
                .save()
                .then(() => this.transitionToRoute('login'))
                .catch(err => alert(err));
        }
    }
});
