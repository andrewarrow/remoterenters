class CreateBuildings < ActiveRecord::Migration[7.0]
  def change
    enable_extension 'pgcrypto'
    create_table :buildings do |t|
      t.uuid :guid, null: false
      t.string :address

      t.timestamps
    end
  end
end
