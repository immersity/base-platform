// app/adapters/application.js
import DS from 'ember-data';
import DataAdapterMixin from 'ember-simple-auth/mixins/data-adapter-mixin';

export default DS.JSONAPIAdapter.extend(DataAdapterMixin, {
    host: 'http://localhost:3000',
    namespace: 'v1',
    authorizer: 'authorizer:application'
});
