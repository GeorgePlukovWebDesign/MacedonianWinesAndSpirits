/**
 * WinesController
 *
 * @description :: Server-side logic for managing wines
 * @help        :: See http://links.sailsjs.org/docs/controllers
 */

module.exports = {
	all: function (req, res) {
    Wines.find().exec(function (err, wines) {
			 res.json(wines);
    });
  },
};
