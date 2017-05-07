/**
 * Wines.js
 *
 * @description :: TODO: You might write a short summary of how this model works and what it represents here.
 * @docs        :: http://sailsjs.org/#!documentation/models
 */

module.exports = {

  attributes: {
    name: {
      type: 'string',
      defaultsTo: 'Unamed'
    },
    description: {
      type: 'string',
      defaultsTo: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'
    },
    type: {
      type:'string',
      enum: ['Red','White', 'Rose','Rakija'],
      defaultsTo: 'Red'
    },
    price: {
      type: 'integer',
      defaultsTo: function() {
        return 0;
      }
    },
    alcoholPercentage: {
      type: 'string',
      defaultsTo: '0.00'
    },
    availability: {
      type: 'string',
      enum: ['available', 'coming-soon,', 'on-its-way', 'out-of-stock', 'not-returning'],
      defaultsTo: 'available'
    },
    year: {
      type: 'string',
      defaultsTo: 'X21X'
    },
    numberPerCase:{
      type: 'integer',
      defaultsTo: 12
    },
    costPerBottle:{
      type: 'float',
      defaultsTo: 0.0
    }
  }
};
