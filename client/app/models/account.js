import Model from 'ember-data/model';
import attr from 'ember-data/attr';

export default Model.extend({
    firstName: attr('string'),
    lastName: attr('string'),
    email: attr('string'),
    password: attr('string'),
    verified: attr('boolean'),
    createdOn: attr('date')
});
