class User < ApplicationRecord
  validates :email, presence: true
  validates :email, format: { with: /[^\s]@[^\s]/, on: :create }
  validates :email, uniqueness: true
  has_and_belongs_to_many :roles
end
