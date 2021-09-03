'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    return queryInterface.createTable("movies", {
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
      duration: {
        allowNull: false,
        type: Sequelize.STRING
    },
      production_house_id: {
          allowNull: false,
          references: {
              model: {
                  tableName: 'production_houses',
              },
              key: 'id'
          },
          type: Sequelize.INTEGER
      },
      release_date: {
        allowNull: false,
        type: Sequelize.DATE
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
    return queryInterface.dropTable("movies");
  }
};
