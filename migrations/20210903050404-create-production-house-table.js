'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    return queryInterface.createTable("production_houses", {
      id: {
          allowNull: false,
          autoIncrement: true,
          primaryKey: true,
          type: Sequelize.INTEGER
      },
      name: {
          allowNull: true,
          unique: true,
          type: Sequelize.STRING
      },
      email: {
          allowNull: false,
          unique: true,
          type: Sequelize.STRING
      },
      location: {
          allowNull:false,
          type: Sequelize.STRING
      },
      created_at: {
          allowNull: false,
          type: Sequelize.DATE
      },
      updated_at: {
          allowNull: false,
          type: Sequelize.DATE
      }
  });
  },

  down: async (queryInterface, Sequelize) => {
    return queryInterface.dropTable("production_houses");
  }
};
