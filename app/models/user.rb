class User < ApplicationRecord
  validates :name, presence: true
  has_and_belongs_to_many :roles
end
